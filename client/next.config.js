/** @type {import('next').NextConfig} */
const path = require('path')

const nextConfig = {
  sassOptions: {
    reactStrictMode: true,
    includePaths: [path.join(__dirname, 'styles')],
  },
}

module.exports = nextConfig
