FROM golang:alpine

# lest be shore where we working
WORKDIR /

RUN apk add --update python ca-certificates unzip

# https://dl.google.com/dl/cloudsdk/channels/rapid/google-cloud-sdk.zip
ENV COMPACT_GOOGLE_CLOUD ./vendor/google-cloud-sdk.zip

# Copy sdk in
ADD $COMPACT_GOOGLE_CLOUD /tmp/google-cloud-sdk.zip

# Unzip it
RUN unzip /tmp/google-cloud-sdk.zip -d / && rm /tmp/google-cloud-sdk.zip

# Copy shared gke install script
COPY install_gke /

# Run it
RUN ./install_gke

ENV PATH /google-cloud-sdk/bin:$PATH

# RUN apk del unzip
# =============================================================================
RUN apk add --update git

COPY . /go/src/github.com/yebo-ecommerce/drone-gke

WORKDIR /go/src/github.com/yebo-ecommerce/drone-gke

RUN go get -d -v

RUN go install -v

# Add the Drone plugin
ADD drone-gke /bin/

ENTRYPOINT ["/bin/drone-gke"]
