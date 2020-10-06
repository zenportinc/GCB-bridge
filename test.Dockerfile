# Run test inside docker
FROM zenportinc/gcb-bridge/test-base
COPY . .
CMD ["bash", "./tests.sh"]