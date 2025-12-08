#!/bin/bash

# Check if version argument is provided
if [ -z "$1" ]; then
    echo "Usage: ./bump_version.sh <new_version_code> [new_version_name]"
    echo "Example: ./bump_version.sh 2 1.0.1"
    exit 1
fi

NEW_CODE="$1"
NEW_NAME="$2"

# If name is not provided, derive it from code (e.g., 2 -> 1.0.2 is hard, so require it for now or keep existing if not provided)
# For simplicity, let's require both or just update code if name missing (but usually name changes too).
if [ -z "$NEW_NAME" ]; then
    echo "Please provide version name as second argument."
    echo "Example: ./bump_version.sh 2 1.0.1"
    exit 1
fi

echo "Bumping version to Code: $NEW_CODE, Name: $NEW_NAME"

# 1. Update frontend/package.json
echo "Updating frontend/package.json..."
sed -i '' "s/\"version\": \".*\"/\"version\": \"$NEW_NAME\"/" frontend/package.json

# 2. Update frontend/android/app/build.gradle
echo "Updating frontend/android/app/build.gradle..."
sed -i '' "s/versionCode .*/versionCode $NEW_CODE/" frontend/android/app/build.gradle
sed -i '' "s/versionName \".*\"/versionName \"$NEW_NAME\"/" frontend/android/app/build.gradle

# 3. Update frontend/src/composables/useAppUpdate.js
echo "Updating frontend/src/composables/useAppUpdate.js..."
sed -i '' "s/code: .*,/code: $NEW_CODE,/" frontend/src/composables/useAppUpdate.js
sed -i '' "s/name: '.*'/name: '$NEW_NAME'/" frontend/src/composables/useAppUpdate.js

# 4. Update backend/config/config.yaml
echo "Updating backend/config/config.yaml..."
sed -i '' "s/version: \".*\"/version: \"$NEW_NAME\"/" backend/config/config.yaml

echo "Version bump complete!"
