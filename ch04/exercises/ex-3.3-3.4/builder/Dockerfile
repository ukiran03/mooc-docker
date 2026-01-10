FROM docker:25-git

WORKDIR /usr/src/app/

# Copy the binary from your current host directory to the container
COPY ./pipeline .

# Ensure the container sees it as an executable
RUN chmod +x ./pipeline

# Use JSON array format for the ENTRYPOINT
ENTRYPOINT ["./pipeline"]
