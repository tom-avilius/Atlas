
package repoinformer


// DetailedRepositoryView shows formatted details from atlas config
func DetailedRepositoryView () bool {

  return listRepositories()
}


func ListRepositories () bool {

  return listOnlyRepositories()
}
