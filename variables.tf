variable cidr {
  type = string
  default = "10.10.0.0/16"
}

variable azs {
  type = list(string)
  default =  ["us-east-2a", "us-east-2b", "us-east-2c"]
}

variable subnets {
  type = list(string)
  default = ["10.10.101.0/24", "10.10.102.0/24", "10.10.103.0/24"]
}