FROM rails:latest

ENV RAILS_VERSION 5.0.0.beta3
ENV RAILS_ENV production

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY . /usr/src/app

RUN bundle install --path vendor/bundle -j4

ENV TZ Asia/Tokyo

CMD ["bin/rails", "server", "-b", "0.0.0.0"]
