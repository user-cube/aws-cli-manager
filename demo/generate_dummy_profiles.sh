#!/bin/bash
# This script generates dummy AWS profile data for screenshot demos
# without revealing real account information

# Create ~/.aws folder if it doesn't exist
mkdir -p ~/.aws

# Create dummy AWS credentials file
cat > ~/.aws/credentials << 'EOF'
[default]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

[demo-dev]
aws_access_key_id = AKIAI44QH8DHBEXAMPLE
aws_secret_access_key = je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY

[demo-staging]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

[demo-prod]
aws_access_key_id = AKIAJ2CZM6EXAMPLE
aws_secret_access_key = OQz0GipLVhEXAMPLEKEY/jE+D8gOJgFu2EXAMPLE

[demo-admin]
aws_access_key_id = AKIAJ34REXAMPLE
aws_secret_access_key = S9P65ERExampleKEY/H7KlExampleKEY
EOF

# Create dummy AWS config file
cat > ~/.aws/config << 'EOF'
[default]
region = us-west-2
output = json

[profile demo-dev]
region = us-east-1
output = json

[profile demo-staging]
region = eu-west-1
output = json

[profile demo-prod]
region = us-west-2
output = json
sso_start_url = https://example.awsapps.com/start
sso_region = us-west-2
sso_account_id = 123456789012
sso_role_name = DemoAdminRole

[profile demo-admin]
region = us-east-2
output = json
sso_start_url = https://example.awsapps.com/start
sso_region = us-east-2
sso_account_id = 987654321098
sso_role_name = DemoAdminRole
EOF

# Set up dummy profiles in aws-cli-manager
mkdir -p ~/.aws/aws-cli-manager

# Create dummy AWS CLI Manager config file
cat > ~/.aws/awsCliManager.yaml << 'EOF'
profiles:
  global:
    region: us-west-2
    ssoEnabled: false
    config: |
      [profile demo-dev]
      region = us-east-1
      output = json

      [profile demo-staging]
      region = eu-west-1
      output = json

      [profile demo-prod]
      region = us-west-2
      output = json
      sso_start_url = https://example.awsapps.com/start
      sso_region = us-west-2
      sso_account_id = 123456789012
      sso_role_name = DemoAdminRole

      [profile demo-admin]
      region = us-east-2
      output = json
      sso_start_url = https://example.awsapps.com/start
      sso_region = us-east-2
      sso_account_id = 987654321098
      sso_role_name = DemoAdminRole
    credentials: |
      [demo-dev]
      aws_access_key_id = AKIAI44QH8DHBEXAMPLE
      aws_secret_access_key = je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY

      [demo-staging]
      aws_access_key_id = AKIAIOSFODNN7EXAMPLE
      aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

      [demo-prod]
      aws_access_key_id = AKIAJ2CZM6EXAMPLE
      aws_secret_access_key = OQz0GipLVhEXAMPLEKEY/jE+D8gOJgFu2EXAMPLE

      [demo-admin]
      aws_access_key_id = AKIAJ34REXAMPLE
      aws_secret_access_key = S9P65ERExampleKEY/H7KlExampleKEY
  demo-dev:
    region: us-east-1
    ssoEnabled: false
    config: |
      region = us-east-1
      output = json
    credentials: |
      aws_access_key_id = AKIAI44QH8DHBEXAMPLE
      aws_secret_access_key = je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
  demo-staging:
    region: eu-west-1
    ssoEnabled: false
    config: |
      region = eu-west-1
      output = json
    credentials: |
      aws_access_key_id = AKIAIOSFODNN7EXAMPLE
      aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
  demo-prod:
    region: us-west-2
    ssoEnabled: true
    config: |
      region = us-west-2
      output = json
      sso_start_url = https://example.awsapps.com/start
      sso_region = us-west-2
      sso_account_id = 123456789012
      sso_role_name = DemoAdminRole
    credentials: |
      aws_access_key_id = AKIAJ2CZM6EXAMPLE
      aws_secret_access_key = OQz0GipLVhEXAMPLEKEY/jE+D8gOJgFu2EXAMPLE
  demo-admin:
    region: us-east-2
    ssoEnabled: true
    config: |
      region = us-east-2
      output = json
      sso_start_url = https://example.awsapps.com/start
      sso_region = us-east-2
      sso_account_id = 987654321098
      sso_role_name = DemoAdminRole
    credentials: |
      aws_access_key_id = AKIAJ34REXAMPLE
      aws_secret_access_key = S9P65ERExampleKEY/H7KlExampleKEY
currentProfile: demo-staging
EOF

chmod +x ~/.aws/aws-cli-manager/dummy_profiles.sh

echo "Dummy profile setup complete!"
echo "Run your demo scripts now to capture screenshots with these dummy profiles."
