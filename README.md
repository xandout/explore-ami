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
xandout$ explore-ami -full
{
  Images: [{
      Architecture: "x86_64",
      BlockDeviceMappings: [{
          DeviceName: "/dev/xvda",
          Ebs: {
            DeleteOnTermination: true,
            Encrypted: false,
            SnapshotId: "snap-0a9551026a7f15871",
            VolumeSize: 8,
            VolumeType: "gp2"
          }
        }],
      CreationDate: "2017-06-23T23:35:49.000Z",
      Description: "Amazon Linux AMI 2017.03.1.20170623 x86_64 HVM GP2",
      EnaSupport: true,
      Hypervisor: "xen",
      ImageId: "ami-a4c7edb2",
      ImageLocation: "amazon/amzn-ami-hvm-2017.03.1.20170623-x86_64-gp2",
      ImageOwnerAlias: "amazon",
      ImageType: "machine",
      Name: "amzn-ami-hvm-2017.03.1.20170623-x86_64-gp2",
      OwnerId: "137112412989",
      Public: true,
      RootDeviceName: "/dev/xvda",
      RootDeviceType: "ebs",
      SriovNetSupport: "simple",
      State: "available",
      VirtualizationType: "hvm"
    }]
}

```

