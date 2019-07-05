workflow "Build dpcmder Project" {
  // We want to run the workflow on each push
  on = "push"

  // Specify which actions should be triggered
  resolves = ["build"]
}

action "build" {
  // My public action
  uses = "./.github-actions"
}
