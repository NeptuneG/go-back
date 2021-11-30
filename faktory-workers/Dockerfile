FROM ruby:3.0.3-alpine3.14

ENV APP_HOME=/app \
    BUNDLE_PATH=/bundle \
    GEM_PATH=/bundle
RUN mkdir $APP_HOME

RUN set -x && \
    echo 'http://dl-cdn.alpinelinux.org/alpine/v3.14/main' >> /etc/apk/repositories && \
    apk add --update --no-cache tzdata ruby-dev make g++ musl-dev && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    gem install bundler --no-document

WORKDIR $APP_HOME

COPY Gemfile ./Gemfile
COPY Gemfile.lock ./Gemfile.lock

RUN bundle config set --local path $BUNDLE_PATH && \
    bundle install

COPY . ./

ENTRYPOINT [ "bundle", "exec", "faktory-worker" ]
CMD [ "-r", "./boot.rb" ]
