# =============================================================================
# Frontend Dockerfile — Next.js (Turborepo + PNPM)
# Multi-stage build untuk optimasi ukuran image dan kecepatan build
# =============================================================================

# ---------------------------------------------------------------------------
# Stage 1: Base — Setup PNPM
# ---------------------------------------------------------------------------
FROM node:24-alpine AS base

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME/bin:$PNPM_HOME:$PATH"

RUN corepack enable && corepack prepare pnpm@11.1.1 --activate
RUN apk add --no-cache libc6-compat

# ---------------------------------------------------------------------------
# Stage 2: Pruner — Prune monorepo untuk app web saja
# ---------------------------------------------------------------------------
FROM base AS pruner

WORKDIR /app

# Install turbo secara global untuk prune
RUN pnpm add -g turbo

# Copy seluruh monorepo
COPY . .

# Prune monorepo — hanya ambil dependensi yang relevan untuk web
RUN turbo prune web --docker

# ---------------------------------------------------------------------------
# Stage 3: Builder — Install dependensi dan build Next.js
# ---------------------------------------------------------------------------
FROM base AS builder

WORKDIR /app

# Copy lockfile dan package.json hasil prune (untuk layer cache optimal)
COPY --from=pruner /app/out/json/ .
COPY --from=pruner /app/out/pnpm-lock.yaml ./pnpm-lock.yaml

# Install dependensi
RUN pnpm install --frozen-lockfile

# Copy source code hasil prune
COPY --from=pruner /app/out/full/ .

# Build aplikasi Next.js via Turborepo
RUN pnpm turbo build --filter=web

# ---------------------------------------------------------------------------
# Stage 4: Runner — Image produksi minimal
# ---------------------------------------------------------------------------
FROM node:24-alpine AS runner

WORKDIR /app

# Tambah non-root user
RUN addgroup --system --gid 1001 nextjs && \
    adduser --system --uid 1001 nextjs

# Copy output standalone Next.js
COPY --from=builder /app/web/.next/standalone ./
COPY --from=builder /app/web/.next/static ./web/.next/static
COPY --from=builder /app/web/public ./web/public

# Atur kepemilikan
RUN chown -R nextjs:nextjs /app

USER nextjs

# Environment variables
ENV NODE_ENV=production
ENV HOSTNAME="0.0.0.0"
ENV PORT=3000

EXPOSE 3000

# Healthcheck
HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
    CMD wget -qO- http://localhost:3000/ || exit 1

# Jalankan Next.js standalone server
CMD ["node", "web/server.js"]
