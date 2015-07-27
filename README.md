# Making immutable images on AWS with Packer and Ansible
Following the [12 Factor App](http://12factor.net/) definitions, we'll put together an AMI that has the software installed on the image, then combine it with configuration for the destination environment for a release AMI.

Think of "build" as installation and release as configuration.  See [build, release, run](http://12factor.net/build-release-run) on the 12 Factor App site.

## Making the build AMI with Packer
The following command will launch an AMI, install the sample web service with a blank configuration file:
```bash
packer build packer/packer-build.json
```

## Add configuration with Packer via Ansible
Once we have an AMI with our application installed, we need to apply the configuration for it.  In this simple example, we're allowing the port the web service attaches to to be configured.  Ansible is used to supply the config values.  We also set the service to start when the instance is spun up.

Supply the AMI ID from the previous step in this command:
```bash
packer build -var 'source_ami=ami-ca056ea2' packer/packer-release.json
```

Now the AMI has the application installed, configuration values are in place and the service will start when the AMI is started.  This makes quick startup for easy autoscaling and eliminates the need to configure new instances as they are spun up.

## Using 'build' and 'release' with build servers
Use Packer's `-machine-readable` flag to get an easily parsable output.  Use this to supply AMI IDs to build steps.

A solid workflow:
 1. Compile source to executable and run unit tests.
 2. Integration tests (optional but recommended).
 3. Package code into a deployable artifact.  May be a standalone executable or many files rolled into an RPM.
 4. 12 Factor 'build' step: using Packer, install the executable or RPM onto an instance, save when done.
 5. 12 Factor 'release' step: using Packer and Ansible, apply desired configuration to the installed application.  Set the application to run on startup then save the AMI.
 6. Deploy new AMI.
