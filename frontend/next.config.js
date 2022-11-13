/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    unoptimized: true, // This because image optimization is not compatible with next export
  }
}

module.exports = nextConfig
