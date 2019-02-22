# Compodoc
This builder can be used to build documentation for your angular projects using a tool called compodoc. If you are new to compodoc [see](https://compodoc.app/) 

## Building this builder
Run the command below to build this builder

```
gcloud builds submit . --config=cloudbuild.yaml
```

## Testing the example
I used the example MVC project from compodoc github to demonstrate the use of this builder. Before you can test this out create a storage bucket in GCP and name it compodoc-demo-todomvc-angularjs, additionally configure the storage bucket for static website hosting. See this [link](https://cloud.google.com/storage/docs/hosting-static-website)

* Switch to examples directory and run command below

```
gcloud builds submit . --config=cloudbuild.yaml
```

