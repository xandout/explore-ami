# View AMI info with Go!!

This is a short example showing how to view AWS EC2 AMI details with Go.

## Usage

```
xandout$ explore-ami
Usage of explore-ami:
  -ami string
        Which AMI to display (default "ami-a4c7edb2")
  -region string
        Which region to query (default "us-east-1")
xandout$ explore-ami
+--------------------------------+---------------------+---------------+---------------+
|          DESCRIPTION           | VIRTUALIZATION TYPE | ROOT DEV NAME | ROOT DEV TYPE |
+--------------------------------+---------------------+---------------+---------------+
| Amazon Linux AMI               | hvm                 | /dev/xvda     | ebs           |
| 2017.03.1.20170623 x86_64 HVM  |                     |               |               |
| GP2                            |                     |               |               |
+--------------------------------+---------------------+---------------+---------------+

```

