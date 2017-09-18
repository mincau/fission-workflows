# Installation

## Setting up Fission (temporary)
As of the moment of writing, the prototype of the Fission Workflow engine has been implemented in Fission using a couple of shortcuts.
In the coming weeks, Fission Workflow will be implemented to fully conform to the Fission Environment API, removing the need for any special modifications to Fission.

To deploy the augmented version of fission, either pull the image:
```bash
docker pull erwinvaneyk/fission-workflow

# To install, update fission.yaml to point to the augmented image: "erwinvaneyk/fission-bundle"
$EDITOR fission.yaml
```

Or, manually build:
```bash
# clone or add a remote to git@github.com:erwinvaneyk/fission.git
git clone git@github.com:erwinvaneyk/fission.git

# Switch to the branch that contains the Fission Workflow integration
git checkout <remote> fission-workflow-integration

# Follow the [guide on compiling Fission](https://github.com/fission/fission/blob/master/Compiling.md)

# Compile and push the fission-bundle to the local Docker repo
(cd fission-bundle/ && bash ./push.sh)
```

After either pulling the custom image or building it yourself, deploy the bundle as specified in [Fission's install guide](http://fission.io/docs/v0.2.1/install/).


## Installing Fission Workflow
Fission Workflow is just another Fission environment.
The environment requires only a single additional property `allowedFunctionsPerContainer` to be set to infinite, to ensure that workflows do not require a workflow environment each.
To deploy the environment run the following:
```bash
kubectl -f https://github.com/fission/fission-workflow/tree/master/build/env/workflow-env.yaml
```

Optionally, you can add the workflow apiserver to the fission deployment. 
This will enable components such as the ui and cli to have a complete view of the workflows and their invocations.
To deploy the apiserver use the following template: 
```bash
kubectl -f https://github.com/fission/fission-workflow/tree/master/build/workflow-apiserver.yaml
```

You're good to go! Check out the [examples](./examples/).