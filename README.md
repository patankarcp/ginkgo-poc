# Ginkgo PoC 🚀

Welcome to the **ginkgo-poc** repository! This project showcases a set of Go-based microservices utilizing Kafka for inter-service communication. We use Kubernetes for deployment, Helm for managing Kubernetes configurations, and Kind for local testing. Here’s an overview of the project:

## 🛠️ What’s Inside

1. **Postgres Microservice**: Manages user data with CRUD operations on a PostgreSQL database.
2. **Redis Microservice**: Caches user data using Redis to improve performance and reduce database load.
3. **Client Microservice**: Provides an interface for creating, updating, deleting, getting, and listing users, interacting with both the Postgres and Redis microservices.
4. **Kafka**: Facilitates communication between the microservices to ensure data consistency and efficient message handling.
5. **Kubernetes**: Manages the deployment and orchestration of the microservices.
6. **Helm**: Helps with templating and managing Kubernetes configurations.
7. **Kind**: Used for testing Kubernetes deployments in a local development environment.

## 🚀 Getting Started

### Prerequisites

- **Go** (version 1.x or later)
- **PostgreSQL** and **Redis** instances
- **Kafka** broker
- **Docker** and **Kind**
- **Kubectl** for interacting with Kubernetes
- **Helm** for managing Kubernetes deployments

## 🚀 Getting Started

Ready to dive in? Follow these steps:

1. **Clone the repository**:
    ```bash
    git clone https://github.com/patankarcp/ginkgo-poc.git
    cd ginkgo-poc
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```
3.  **Run tests**:
    
    **TODO NOTE**: 
        - Add test instructions/commands below.
    ```bash
    
    ```

For detailed instructions, check out the [README](./README.md).

## 💡 Want to Contribute?

We’d love your help! If you're interested in contributing, please see our [Contributing Guide](./CONTRIBUTING.md) to get started.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more information.

## 🔗 Explore More

- Check out the [documentation](./docs) for more details.
- If you have any questions or ideas, feel free to [open an issue](https://github.com/patankarcp/ginkgo-poc/issues).

Happy coding! 🎉
