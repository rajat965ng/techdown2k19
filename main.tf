
provider "example" {
  address = "http://localhost"
  port = 3001
  token = "s3cr3t"
}

resource "example_item" "test" {
  name = "Ravikant"
  description = "this is an item"
  tags = ["hello", "world"]
}