name      = "cos"
implement = ["prefix_lister", "dir_lister"]

service {

  new {
    required = ["credential"]
  }

  op "create" {
    required = ["location"]
  }
  op "delete" {
    required = ["location"]
  }
  op "get" {
    required = ["location"]
  }
  op "list" {
    required = ["storager_func"]
  }
}

storage {

  new {
    required = ["location", "name"]
    optional = ["work_dir"]
  }

  op "list_dir" {
    optional = ["dir_func", "file_func"]
  }
  op "list_prefix" {
    required = ["object_func"]
  }
  op "write" {
    required = ["size"]
    optional = ["checksum", "storage_class"]
  }
}

pairs {

  pair "storage_class" {
    type = "string"
  }
}

infos {

  info "object" "meta" "storage-class" {
    type = "string"
  }
}
