data "aws_security_group" "_All_Open" {
    name = "_All_Open"
}

resource "aws_instance" "Go-App" {
  ami           = "ami-00b8917ae86a424c9"
  instance_type = "t2.micro"


  user_data = file("../scripts/user_data.sh")

  tags = {
    Name = var.instance_name
  }

    vpc_security_group_ids = [data.aws_security_group._All_Open.id]  
}