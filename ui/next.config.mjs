/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "standalone",
  webpack: (config, { dev, isServer }) => {
    // Enable source maps explicitly
    if (dev && !isServer) {
      config.devtool = "source-map";
    }

    return config;
  },
};

export default nextConfig;
