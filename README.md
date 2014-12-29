# Making immutabile images on AWS with Packer and Ansible

## Put an application on an AMI with Packer
```bash
packer build packer/packer-application.json
```

## Add configuration values to the image with Ansible
Done in previous step for now

## Deploy the image
Launch the AMI packer made
