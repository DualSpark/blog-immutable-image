# Making immutabile images on AWS with Packer and Ansible
Following the 12 Factor Application definitions, we'll put together an AMI that has the software installed on the image, then combine it with configuration for the destination environment for a release AMI.

## Making the build AMI with Packer
The following command will launch an AMI, install the sample web service with a blank configuration file:
```bash
packer build packer/packer-build.json
```

## Add configuration with Packer via Ansible
Once we have an AMI with our application installed, we need to apply the configuration for it.  In this simple example, we're allowing the port the web service attaches to to be configured.  Ansible is used to supply the config values.  We also set the service to start by default by using the `chkconfig goapp on` command.

Supply the AMI ID from the previous step in this command:
```bash
packer build -var 'source_ami=ami-ca056ea2' packer/packer-release.json
```

Now the AMI has the application installed, configuration values are in place and the service will start when the AMI is started.  This makes quick startup for easy autoscaling and eliminates the need to configure new instances as they are spun up.
