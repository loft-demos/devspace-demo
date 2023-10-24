# devspace-demo

pipelines:
  # Override the default pipeline for 'devspace dev'
  dev: |-
    kubectl delete src api -n api
    run_dependencies --all
    ensure_pull_secrets --all
    build_images --all
    create_deployments --all --exclude api-staging
    start_dev --all