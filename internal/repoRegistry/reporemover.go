
package reporegistry


func removeRepository (repoName string,  onlyRemoveReference bool) bool {

  if (onlyRemoveReference) {

    return deleteYaml("/home/tom-avilius/Projects/Atlas/tests/testConfig.yaml", repoName)
  }

  // TODO: Implement a system to delete local folders too. For that,
  // we would need to decide a system to store all folders in a single place. So we can look up by repoName.
  // However, we might also store the absolute repo path in the config itself and use a function like searchYaml
  // to retrieve the path.

  return true
}
