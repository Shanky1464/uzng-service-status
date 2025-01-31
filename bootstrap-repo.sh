#!/bin/bash

set -o pipefail
set +o errexit

export LC_ALL=C

readonly TEMPLATE_NAME='hello-world-go'
readonly SERVICE_NAME_PATTERN='^[a-zA-Z]+[a-zA-Z0-9-]*[^-]$'
readonly CICD_PATH='.github/workflows/cicd.yml'
readonly README_PATH='README.md'
readonly CODEOWNERS_PATH='CODEOWNERS'
readonly DD_SERVICE_CATALOG_PATH='service.datadog.yaml'

function die {
  printf "[%s]: $*\n" "$(date +'%Y-%m-%dT%H:%M:%S%z')" >&2
  exit 1
}

command -v git > /dev/null || die "Application \033[1mgit\033[0m is required but not installed. Install git to proceed."
command -v yq > /dev/null || die "Application \033[1myq\033[0m is required but not installed. Install yq to proceed."
command -v sed > /dev/null || die "Application \033[1msed\033[0m is required but not installed. Install sed to proceed."

# Evaluate the new service name
repo_name="$(basename "$(git rev-parse --show-toplevel)")"
if [[ $# == 1 ]]; then
  repo_name="$1"
fi

if [[ ! "$repo_name" =~ $SERVICE_NAME_PATTERN ]]; then
  die "Service name is invalid"
fi


# Continue only if in the default Git branch
default_branch="$(basename "$(git symbolic-ref refs/remotes/origin/HEAD)")"
current_branch="$(git rev-parse --abbrev-ref HEAD)"
bootstrap_branch="bootstrap-repo"

if [[ "$current_branch" != "$default_branch" ]]; then
  die "You must check out the default branch: $default_branch"
fi

git switch --create "$bootstrap_branch"


# Detect the installed version of sed
sed --version > /dev/null 2>&1
case $? in
  0) # GNU sed
    function sed_i { sed -i "$@"; }
    ;;
  1) # BSD sed
    function sed_i { sed -i '' "$@"; }
    ;;
  127)
    die "Application \033[1msed\033[0m is required but not installed. Install sed to proceed."
    ;;
  *)
    die "Something went wrong.."
    ;;
esac


# Rename the new service
function rename {
  # shellcheck disable=SC2046
  sed_i "$@" $(find . -type f -not -path "./.git/*" -not -path "./.garden/*")
}

rename "s/platform-${TEMPLATE_NAME}/${repo_name}/g"
rename "s/${TEMPLATE_NAME}/${repo_name}/g"

template_alt_name="${TEMPLATE_NAME//-/_}"
service_alt_name="${repo_name//-/_}"
rename "s/platform_${template_alt_name}/${service_alt_name}/g"
rename "s/${template_alt_name}/${service_alt_name}/g"

find . -depth -type d -name "platform-${TEMPLATE_NAME}" -execdir git mv {} "${repo_name}" \;
find . -depth -type d -name "${TEMPLATE_NAME}" -execdir git mv {} "${repo_name}" \;
find . -depth -type d -name "platform_${template_alt_name}" -execdir git mv {} "${service_alt_name}" \;
find . -depth -type d -name "${template_alt_name}" -execdir git mv {} "${service_alt_name}" \;

printf "The service name has been renamed in the following files:\n"
git status --short

git commit --all --message="Rename service name to \`$repo_name\`"


# Update README.md
[[ -e "$README_PATH" ]] || die "$README_PATH file not found"

sed_i "s#^A template for platform services written in.*#\
This service has been created from a template:\n\
[platform-$TEMPLATE_NAME](https://github.com/uniphore/platform-$TEMPLATE_NAME)#" "$README_PATH"

delete_start='^## Description'
delete_stop='^## Requirements'
sed_i -e "/$delete_start/,/$delete_stop/{" -e "/$delete_stop/!d"  -e "}" "$README_PATH"

# shellcheck disable=SC2016
sed_i '$d; $d' "$README_PATH"

git add "$README_PATH"
git commit --message="Update \`README.md\`"


# Enable CI/CD pipeline
[[ -e "$CICD_PATH" ]] || die "The CI/CD connector not found at $CICD_PATH"

yq -i 'del(.jobs.CI.if)' "$CICD_PATH"

# Only stage the deleted attribute
cicd_diff="$(git --no-pager diff "$CICD_PATH")"
git restore "$CICD_PATH"
HUNK_START='^@@ [-+, [:digit:]]\{1,\}@@.*$'
cicd_diff="$(sed -ne "/$HUNK_START/{" -e x -e 's/.*/&x/' -e '/.\{2\}/q' -e 'x' -e '}' -e p <<< "$cicd_diff")"
git apply <<< "$cicd_diff"

git add "$CICD_PATH"
git commit --message="Enable CI/CD pipeline"


# Update DataDog Service Catalog and CODEOWNERS
PS3="Select your team (1-8): "
options=(
  "AI"
  "GTS"
  "Interact"
  "Platform"
  "U-Analyze NG"
  "U-Assist"
  "U-Capture NG"
  "U-SelfServe"
  "Other"
)

select opt in "${options[@]}"
do
  case $opt in
    "AI")
      team_id=ai
      codeowners=datascience-write
      break;;
    "GTS")
      team_id=global-technical-solutions
      codeowners=delivery
      break;;
    "Interact")
      team_id=xconsole
      codeowners=interact-code-owners
      break;;
    "Platform")
      team_id=platform
      codeowners=x-platform
      break;;
    "U-Analyze NG")
      team_id=uanalyze-ng
      codeowners=u-analyze-approvers
      break;;
    "U-Assist")
      team_id=uassist
      codeowners=u-assist-maintainer
      break;;
    "U-Capture NG")
      team_id=ucapture-ng
      codeowners=ucaptureng
      break;;
    "U-SelfServe")
      team_id=uselfserve
      codeowners=u-selfserve-write
      break;;
    "Other")
      read -r -p "Enter ID of the team (or contact #platform-support): " team_id
      [[ -n "$team_id" ]] && break
      ;;
    *) echo "Invalid choice. Please select a number from 1 to 8.";;
  esac
done

export team_id
yq -i '.team = env(team_id)' "$DD_SERVICE_CATALOG_PATH"
sed_i "s/x-platform/${codeowners}/" "$CODEOWNERS_PATH"
unset team_id

git add "$DD_SERVICE_CATALOG_PATH" "$CODEOWNERS_PATH"
git commit --message="Update service ownership"

# Delete the bootstrap script
git rm "$0"
git commit --message="Delete bootstrap script"


# Summary
git log "$default_branch.."

printf "\n\n"
printf "Please run the following command in order to push the changes to GitHub and open Pull Request\n"
printf "*********************************************************************************************\n"
printf "****                    git push --set-upstream origin %s                    ****\n" "$bootstrap_branch"
printf "*********************************************************************************************\n"

exit 0
