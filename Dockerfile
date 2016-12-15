FROM alpine

# https://dl.google.com/dl/cloudsdk/channels/rapid/google-cloud-sdk.zip
ENV COMPACT_GOOGLE_CLOUD ./vendor/google-cloud-sdk.zip

RUN apk add --update python ca-certificates unzip
# ============================================================================
ENV CLOUDSDK_PYTHON_SITEPACKAGES 1
# Install the Google Cloud SDK.
ADD $COMPACT_GOOGLE_CLOUD /tmp/google-cloud-sdk.zip
RUN unzip /tmp/google-cloud-sdk.zip -d / && rm /tmp/google-cloud-sdk.zip

# Install kubectl
RUN google-cloud-sdk/install.sh --quiet
RUN google-cloud-sdk/bin/gcloud components install kubectl

# Clean up
RUN rm -rf google-cloud-sdk/.install

ENV PATH /google-cloud-sdk/bin:$PATH

RUN apk del unzip

# Add the Drone plugin
ADD drone-gke /bin/

ENTRYPOINT ["/bin/drone-gke"]
