$projectID = gcloud config get-value project;

[Net.ServicePointManager]::SecurityProtocol = "tls12, tls11, tls";
$ProgressPreference = 'SilentlyContinue';

echo "Downloading Git 2.18";
Set-Variable -name GIT_URL -value "https://github.com/git-for-windows/git/releases/download/v2.18.0.windows.1/MinGit-2.18.0-64-bit.zip";
Invoke-WebRequest -Uri $GIT_URL -OutFile git.zip;

echo "Downloading Go 1.10";
Set-Variable -name GO_URL -value "https://dl.google.com/go/go1.10.3.windows-386.zip";
Invoke-WebRequest -Uri $GO_URL -Outfile go.zip;

gcloud --quiet auth configure-docker;

echo "Running Docker build";
docker build -t gcr.io/$projectID/go-windows .;
if ($?) {
    docker push gcr.io/$projectID/go-windows;
}
