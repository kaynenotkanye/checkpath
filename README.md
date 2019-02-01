Simple GO app to check if file or path exists. Panics if specified path does not exist. As a use case, this app was created for two reasons:
1. As a tool for unit testing of Packer AMI builds (for Amazon EC2s).
2. Just me learning Go, kind of "Hello Worldish"

### Example usage
```sh
checkpath -p /path/to/war/file.war
checkpath -p /etc/tomcat8
```

Referencing the above commands, if used during a Packer build, it is designed to fail the Packer build if the `file.war` does not exist. It will also fail the Packer build if the `/etc/tomcat8` directory does not exist.  The app also does a dual write to console as well as a logfile. If Packer is run through orchestration/automation such as Jenkins or a Gitlab runner, you would be able to see the exact failure message in a console.  The logfile itself is most likely unnecessary, but the logfile can be used for validation purposes after launching EC2s from the AMI.
