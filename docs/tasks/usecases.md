# Use Cases

## Deploying the head of the `main` branch

The general way of deployment is deploying the head of the `main` branch that has only verified commits. To constrain deploying only the head of the `main` branch, we can use the `auto_merge` parameter of GitHub [deployment API](https://docs.github.com/en/rest/reference/repos#create-a-deployment) to ensure that the deployment is the head of the branch, and set `deployable_ref` field with the `main`.

```yaml
envs:
  - name: production
    auto_merge: true
    deployable_ref: main
    production_environment: true
```

## Deploying with the version

The versioning is the general way to specify the artifact or the commit, and GitHub provides the release page for versioning. If your team (or organization) wants to constrain deploying with the version, you can use the `deployable_ref` field like below.

```yaml
envs:
  - name: production
    auto_merge: false
    deployable_ref: 'v.*\..*\..*'       # Semantic versioning
    production_environment: true
```

## Deploying the artifact

The artifact could be a binary file from compiling source codes or a docker image, which means we have to build the artifact before we deploy. The commit status is the best way to verify if the artifact exists or not. The builder, such as GitHub Action or Circle CI, posts the commit status after building the artifact, and we can verify it by the `required_contexts` parameter when we deploy. You can reference the `deploy.yml` of Gitploy.

```yaml
envs:
  - name: production
    auto_merge: false
    required_contexts:
      - "publish-image"         # The commit status of building the artifact.
    production_environment: true
```

## Questions?

We are always happy to help with questions you might have. You can post questions or comments to our [community](https://github.com/gitploy-io/gitploy/discussions). 

And we are very welcome to share your deployment pipeline at our [community](https://github.com/gitploy-io/gitploy/discussions).