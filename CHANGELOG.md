# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v1.3.0]
### Added
- Added FidelityFX Super Res
  - AMD FidelityFX Super Resolution shader
- Added AMD FidelityFX Super Resolution shader with Subpixel Morphological Antialiasing
  - [What is Subpixel Morphological Antialiasing?](http://www.iryoku.com/smaa/)
- Added Technicolor film shader

### Changes
- Improve error handling in `build.go`.
- Updated CRT Dr Venom Fast
- Updated CRT Fast
- Removed unused images from Gameboy shader

## [v1.2.0]
### Added

- Start using CHANGELOG.md

### Changed

- Updated CRT New Pixie, which includes `Rolling Scanlines` parameter.
- Updated VT220, with many improvements.
- Updated CRT Fast, with improvements for rotated games.
- Updated Go dependency to version 1.17.

[Unreleased]: https://github.com/OpenEmu/slang-shaders/compare/v1.3.0...HEAD
[v1.3.0]: https://github.com/OpenEmu/slang-shaders/compare/v1.2.0...v1.3.0
[v1.2.0]: https://github.com/OpenEmu/slang-shaders/compare/v1.1...v1.2.0