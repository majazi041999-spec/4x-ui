# 🤖 Phantom Protocol - AI Execution Guide

## 📋 Overview

This is a **complete, AI-executable implementation guide** for building Phantom Protocol - a next-generation anti-censorship protocol based on Xray-core.

**For AI Agent**: You are tasked with implementing this protocol by following the instructions in the accompanying Word document (`Phantom_Protocol_AI_Executable_Guide.docx`) **sequentially and precisely**.

---

## 🎯 Your Mission

Build Phantom Protocol from scratch by:
1. Following each task in sequential order
2. Executing all commands exactly as specified
3. Validating each step before proceeding
4. Creating all specified files with exact code
5. Running all tests and ensuring they pass
6. Committing code with proper git workflow

---

## 📁 Required Files

Before starting, ensure you have:
- [ ] `Phantom_Protocol_AI_Executable_Guide.docx` - Main implementation guide
- [ ] This README.md file
- [ ] Access to GitHub account
- [ ] Ubuntu 22.04+ environment (or Docker)

---

## 🚀 Quick Start

### Step 0: Environment Setup

**CRITICAL**: Execute ALL prerequisites before starting development.

```bash
# 1. Verify system
uname -a  # Should be: Linux, Ubuntu 22.04+
nproc     # Recommended: 4+ cores
free -h   # Recommended: 8+ GB RAM

# 2. Install tools (execute from the guide: Section 0.1.1)
# Go 1.21+, Rust 1.75+, Node.js 18+, Docker, BPF tools

# 3. Setup repository (execute from the guide: Section 0.2)
git clone https://github.com/[YOUR_USERNAME]/phantom-protocol.git
cd phantom-protocol

# 4. Verify setup
go version    # >= 1.21
rustc --version  # >= 1.75
node --version   # >= 18
docker --version
bpftool version
```

---

## 📖 How to Use This Guide

### For AI Agents:

1. **Read Document Sequentially**
   - Open `Phantom_Protocol_AI_Executable_Guide.docx`
   - Start from PART 0: Prerequisites
   - Execute each section in order

2. **Task Structure**
   Each task contains:
   - 📝 Objective: What you're building
   - 📍 Location: Where files go
   - 💻 Code: Exact code to write
   - ✅ Validation: Tests to run
   - ☑️ Checklist: Items to verify

3. **Execution Pattern**
   ```
   For each Task:
     1. Read objective
     2. Create specified files
     3. Write exact code provided
     4. Run validation commands
     5. Verify checklist items
     6. Commit with specified message
     7. Move to next task
   ```

4. **Validation is Mandatory**
   - After each task, run ALL validation steps
   - If tests fail, debug before proceeding
   - All checklist items must be ✅ before next task

5. **Git Workflow**
   ```bash
   # For each task/week:
   git checkout -b feature/phase1-week1
   # ... do work ...
   git add .
   git commit -m "feat(component): description"
   git push origin feature/phase1-week1
   # Create PR, review, merge to develop
   ```

---

## 📊 Project Phases

### Phase 1: Core Protocol (Weeks 1-12) ⏱️ 2-3 months
- QUIC Manager with ECH
- Connection Migration
- Fragment Engine
- Basic Mimicry (3 profiles)
- Integration & Testing

**Deliverable**: Working prototype with anti-detection

### Phase 2: Performance (Weeks 13-20) ⏱️ 1-2 months
- eBPF/XDP Implementation
- io_uring Zero-Copy
- BBR v3 Congestion Control
- Multiplexing & QoS

**Deliverable**: 1.5-2x performance improvement

### Phase 3: Anti-Detection (Weeks 21-28) ⏱️ 1-2 months
- 10+ Traffic Profiles
- ML Profile Selection
- Dynamic Morphing
- Timing Obfuscation

**Deliverable**: <1% detection rate after 90 days

### Phase 4: Integration (Weeks 29-44) ⏱️ 2-3 months
- Xray-core Integration
- 3x-ui Panel Integration
- Multi-Platform Clients
- Testing & Hardening
- Documentation & Release

**Deliverable**: Production-ready v1.0.0

---

## 🔍 Task Format Example

Every task in the guide follows this format:

```
Task X.Y: [Task Name]

📝 Objective: [What you're building]

📍 Location: [File paths]

💻 Implementation:
[Exact code blocks to write]

✅ Validation Steps:
[Commands to run]

☑️ Checklist:
□ Item 1
□ Item 2
...

🔀 Git:
git commit -m "feat(component): description"
```

---

## 📂 Project Structure

After completion, your repository should look like:

```
phantom-protocol/
├── core/                    # Core protocol engine
│   ├── mimicry/            # Mimicry engine
│   ├── quic/               # QUIC manager
│   ├── fragment/           # Fragment engine
│   ├── timing/             # Timing obfuscator
│   └── transport/          # Zero-copy transport
├── server/                  # Server application
├── client/                  # Client library
├── ebpf/                    # eBPF programs
├── profiles/                # Traffic profiles
├── tests/                   # Tests
├── docs/                    # Documentation
├── scripts/                 # Build/deploy scripts
└── deployments/             # Docker/K8s configs
```

---

## 🧪 Testing Strategy

Execute tests at every level:

```bash
# Unit tests (after each file)
go test ./core/quic -v -cover

# Integration tests (after each week)
go test ./tests/integration -v

# E2E tests (after each phase)
go test ./tests/e2e -v

# Performance tests (Phase 2)
./scripts/benchmark.sh

# Security tests (Phase 3)
./scripts/security-test.sh
```

**Rule**: No task is complete until tests pass.

---

## 🎓 Implementation Guidelines

### Code Quality Standards

1. **Follow Exact Specifications**
   - Code in the guide is battle-tested
   - Don't deviate unless you have good reason
   - If you improve something, document why

2. **Comments**
   - Add comments for complex logic
   - Use TODO for future improvements
   - Reference GitHub issues when relevant

3. **Error Handling**
   ```go
   // ✅ GOOD
   if err != nil {
       return fmt.Errorf("failed to do X: %w", err)
   }
   
   // ❌ BAD
   if err != nil {
       panic(err)  // Never panic in library code
   }
   ```

4. **Testing**
   - Minimum 80% code coverage
   - Test happy path and error cases
   - Use table-driven tests

### Git Commit Messages

```bash
# Format: <type>(<scope>): <description>

feat(quic): implement ECH encryption
fix(fragment): correct TLS record header
test(mimicry): add profile selection tests
docs(api): update QUIC manager documentation
chore(deps): update quic-go to v0.40.1
```

Types: `feat`, `fix`, `docs`, `test`, `chore`, `refactor`, `perf`

---

## 🔧 Integration with Xray-core & 3x-ui

### Xray-core Integration (Phase 4, Week 29-31)

**Location**: `xray-core-fork/transport/internet/phantom/`

**Key Files**:
- `config.proto` - Protocol buffer definition
- `config.pb.go` - Generated protobuf code
- `phantom.go` - Main transport implementation
- `dialer.go` - Dialer registration

**Steps**:
1. Fork https://github.com/XTLS/Xray-core
2. Create phantom transport directory
3. Implement transport interface
4. Register in Xray initialization
5. Build and test

### 3x-ui Integration (Phase 4, Week 32-34)

**Location**: `3x-ui-fork/web/`

**Key Files**:
- `service/phantom.go` - Backend service
- `controller/phantom.go` - API handlers
- `html/xui/inbound_modal.html` - UI components
- `assets/js/inbound.js` - Frontend logic

**Steps**:
1. Fork https://github.com/MHSanaei/3x-ui
2. Add Phantom protocol service
3. Create UI components
4. Add API endpoints
5. Test in panel

---

## 📈 Progress Tracking

### Use the Complete Project Checklist

At the end of the guide document, there's a comprehensive checklist:
- [ ] Phase 1 items (12 weeks)
- [ ] Phase 2 items (8 weeks)
- [ ] Phase 3 items (8 weeks)
- [ ] Phase 4 items (16 weeks)
- [ ] Final Deliverables (12 items)

**Update this checklist as you complete each task.**

### Performance Benchmarks

Track these metrics throughout development:

| Metric | Phase 1 Target | Phase 2 Target | Final Target |
|--------|---------------|---------------|--------------|
| Throughput | 120 Mbps | 180 Mbps | 200 Mbps |
| Latency | 40-45ms | 30-35ms | 25-30ms |
| CPU Usage | 70% | 50% | 40% |
| Detection Rate | 10% | 3% | <1% |

---

## 🐛 Debugging Tips

### Common Issues

1. **Build Failures**
   ```bash
   # Clear cache and rebuild
   go clean -cache -modcache -i -r
   rm -rf vendor/
   go mod tidy
   go build -v
   ```

2. **Test Failures**
   ```bash
   # Run with verbose output
   go test -v -race -cover ./...
   
   # Run specific test
   go test -v -run TestSpecificFunction
   
   # Show test output
   go test -v -count=1  # Disable cache
   ```

3. **eBPF Issues**
   ```bash
   # Check kernel version
   uname -r  # Need 5.15+
   
   # Verify BPF support
   bpftool feature
   
   # Load manually
   sudo bpftool prog load phantom_xdp.o /sys/fs/bpf/phantom
   ```

### Getting Help

If stuck on a task:
1. Re-read the task objective
2. Check validation steps - what's failing?
3. Review similar code in the codebase
4. Check referenced GitHub issues/PRs
5. Test individual components in isolation

---

## 📦 Final Deliverables

When all phases are complete, you should have:

### Repositories
- [x] phantom-protocol (main implementation)
- [x] Xray-core fork (with Phantom transport)
- [x] 3x-ui fork (with Phantom support)
- [x] v2rayNG fork (Android client)
- [x] v2rayN fork (Windows client)

### Documentation
- [x] README.md (project overview)
- [x] ARCHITECTURE.md (technical details)
- [x] API.md (API documentation)
- [x] DEPLOYMENT.md (installation guide)
- [x] CONTRIBUTING.md (contribution guide)

### Releases
- [x] v1.0.0-beta.1 (beta testing)
- [x] v1.0.0 (stable release)

### Tests & Benchmarks
- [x] Unit tests (>80% coverage)
- [x] Integration tests
- [x] E2E tests
- [x] Performance benchmarks
- [x] Security audit report

---

## ⚠️ CRITICAL REMINDERS

### For AI Agents:

1. **NEVER SKIP VALIDATION**
   - Every task has validation steps
   - Run them ALL
   - Don't proceed if they fail

2. **FOLLOW CODE EXACTLY**
   - Code in guide is tested
   - Exact variable names matter
   - Import paths must be correct

3. **SEQUENTIAL EXECUTION MANDATORY**
   - Tasks build on each other
   - Skipping breaks everything
   - Dependencies must exist

4. **GIT WORKFLOW IS CRITICAL**
   - Commit after each task
   - Use proper messages
   - Push to feature branches

5. **TESTING IS NOT OPTIONAL**
   - No untested code
   - Fix failures immediately
   - Maintain coverage >80%

---

## 🎯 Success Criteria

Project is complete when:

✅ All 44 weeks of tasks are done
✅ All checklist items are marked complete
✅ All tests pass (unit, integration, e2e)
✅ Performance targets are met
✅ Security audit passes
✅ Xray-core integration works
✅ 3x-ui panel works
✅ Clients can connect successfully
✅ Detection rate <1% after 90 days
✅ Documentation is complete
✅ v1.0.0 is released

---

## 📞 Support

- **Documentation**: See `Phantom_Protocol_AI_Executable_Guide.docx`
- **Issues**: Create detailed GitHub issues
- **Questions**: Reference specific task numbers

---

## 📄 License

This implementation guide is provided as-is for educational and development purposes.

Phantom Protocol itself will be released under MIT License upon completion.

---

**Version**: 1.0  
**Date**: February 2026  
**For**: AI-Assisted Development

---

# 🚀 BEGIN EXECUTION

Start with:
1. Open `Phantom_Protocol_AI_Executable_Guide.docx`
2. Go to PART 0: Prerequisites
3. Execute Section 0.1.1: Required Tools Installation
4. Continue sequentially

**Good luck! 🎉**
