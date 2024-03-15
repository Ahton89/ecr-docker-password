# ECR-DOCKER-PASSWORD

A small github action for obtaining the password from the AWS ECR for using it as a password for docker actions.

## Example usage
```yaml
    - name: Get ECR password
      uses: ahton89/ecr-docker-password@v0.0.9
      with:
        access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
        secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
```

## Inputs
```yaml
  access-key-id:
    description: 'AWS Access Key ID'
    required: true
  secret-access-key:
    description: 'AWS Secret Access Key'
    required: true
  region:
    description: 'AWS region'
    required: false
    default: 'us-east-2'
```

## Outputs
```yaml
  ecr-docker-password:
    description: 'The password for the ECR'
```

Use the output in the docker login action
```yaml
    - name: Login to Amazon ECR
      uses: docker/login-action@v1
      with:
        registry: ${{ secrets.AWS_ACCOUNT_ID }}.dkr.ecr.${{ secrets.AWS_REGION }}.amazonaws.com
        username: AWS
        password: ${{ steps.get-ecr-password.outputs.ecr-docker-password }}
```

No warranty is given, use at your own risk. 🤷‍
