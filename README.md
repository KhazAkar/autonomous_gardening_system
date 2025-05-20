# Autonomous Gardening System

Full stack ( LaFrite FE + BE and web FE+BE) implementation of autonomous gardening system, which will help my bell peppers grow 🌱

## 🤝 Contributing - Device

0. Install uv

1. Clone the repo

```bash
git clone https://codeberg.org/KhazAkar/autonomous_gardening_system
cd autonomous_gardening_system/device
```

2. Build the project

```bash
uv sync --all-extras
docker build -t local_device
```

4. Run the tests

```bash
uv pytest tests/
```

If you'd like to contribute, please fork the repository and open a pull request to the `master` branch.
