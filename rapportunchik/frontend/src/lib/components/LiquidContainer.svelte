<script lang="ts">
  import { onMount } from 'svelte';
  import { appState } from '../state.svelte';
  import { i18n } from '../i18n';
  import { Trophy, Sparkles } from 'lucide-svelte';

  let canvasEl: HTMLCanvasElement;
  let animId: number;

  // Visual state
  let currentLevel = 0; // Starts from 0 and smoothly fills up
  let phase = 0;

  interface Bubble {
    x: number;
    y: number;
    radius: number;
    speed: number;
    opacity: number;
  }

  let bubbles: Bubble[] = [];

  function initBubbles(w: number, h: number) {
    bubbles = [];
    for (let i = 0; i < 22; i++) {
      bubbles.push({
        x: Math.random() * w,
        y: h - Math.random() * (h * 0.4),
        radius: 1.5 + Math.random() * 3,
        speed: 0.6 + Math.random() * 1.2,
        opacity: 0.2 + Math.random() * 0.5
      });
    }
  }

  onMount(() => {
    const ctx = canvasEl.getContext('2d');
    if (!ctx) return;

    const resize = () => {
      const rect = canvasEl.getBoundingClientRect();
      const dpr = window.devicePixelRatio || 1;
      canvasEl.width = rect.width * dpr;
      canvasEl.height = rect.height * dpr;
      ctx.scale(dpr, dpr);
      initBubbles(rect.width, rect.height);
    };

    resize();
    window.addEventListener('resize', resize);

    const render = () => {
      const rect = canvasEl.getBoundingClientRect();
      const w = rect.width;
      const h = rect.height;

      // Calculate target fill percentage [0, 1]
      const targetMoney = appState.overallStats?.revenueTarget || 20000;
      const earned = appState.overallStats?.totalMoneyEarned || 0;
      const targetPercent = Math.min(1, Math.max(0.04, earned / targetMoney));

      // Slow organic fill animation to target level
      currentLevel += (targetPercent - currentLevel) * 0.035;

      ctx.clearRect(0, 0, w, h);

      const waterHeight = h * currentLevel;
      const baseY = h - waterHeight;

      phase += 0.035;

      // 1. Draw Back Wave (Darker, slightly out of phase)
      ctx.save();
      ctx.beginPath();
      ctx.moveTo(0, h);
      for (let x = 0; x <= w; x += 4) {
        const y = baseY + Math.sin(x * 0.015 + phase * 1.3) * 9 + Math.cos(x * 0.02 - phase) * 4;
        ctx.lineTo(x, y);
      }
      ctx.lineTo(w, h);
      ctx.closePath();
      const gradBack = ctx.createLinearGradient(0, baseY - 10, 0, h);
      gradBack.addColorStop(0, 'hsla(234, 54%, 28%, 0.7)');
      gradBack.addColorStop(1, 'hsla(234, 54%, 14%, 0.9)');
      ctx.fillStyle = gradBack;
      ctx.fill();
      ctx.restore();

      // 2. Draw Front Wave (Vibrant indigo with highlight crest)
      ctx.save();
      ctx.beginPath();
      ctx.moveTo(0, h);
      for (let x = 0; x <= w; x += 4) {
        const y = baseY + Math.sin(x * 0.018 - phase) * 11 + Math.sin(x * 0.032 + phase * 0.8) * 5;
        ctx.lineTo(x, y);
      }
      ctx.lineTo(w, h);
      ctx.closePath();

      const gradFront = ctx.createLinearGradient(0, baseY - 15, 0, h);
      gradFront.addColorStop(0, 'hsla(234, 100%, 62%, 0.88)');
      gradFront.addColorStop(0.35, 'hsla(234, 85%, 45%, 0.92)');
      gradFront.addColorStop(1, 'hsla(234, 54%, 20%, 0.98)');
      ctx.fillStyle = gradFront;
      ctx.fill();

      // Glowing Crest Line
      ctx.lineWidth = 3;
      ctx.strokeStyle = 'hsla(234, 100%, 87%, 0.85)';
      ctx.shadowColor = 'hsl(234, 100%, 60%)';
      ctx.shadowBlur = 12;
      ctx.stroke();
      ctx.restore();

      // 3. Draw Organic Rising Bubbles
      ctx.save();
      for (const b of bubbles) {
        b.y -= b.speed;
        b.x += Math.sin(b.y * 0.05 + phase) * 0.4;
        if (b.y < baseY) {
          b.y = h + Math.random() * 20;
          b.x = Math.random() * w;
        }

        ctx.beginPath();
        ctx.arc(b.x, b.y, b.radius, 0, Math.PI * 2);
        ctx.fillStyle = `hsla(234, 100%, 87%, ${b.opacity})`;
        ctx.fill();
      }
      ctx.restore();

      animId = requestAnimationFrame(render);
    };

    render();

    return () => {
      window.removeEventListener('resize', resize);
      cancelAnimationFrame(animId);
    };
  });
</script>

<div class="liquid-wrapper">
  <div class="liquid-container">
    <!-- Milestone ticks on side -->
    <div class="milestones-bar">
      <div class="milestone-mark" style="top: 5%;"><span>20k €</span></div>
      <div class="milestone-mark" style="top: 25%;"><span>15k €</span></div>
      <div class="milestone-mark" style="top: 50%;"><span>10k €</span></div>
      <div class="milestone-mark" style="top: 75%;"><span>5k €</span></div>
    </div>

    <!-- Canvas for wave animation -->
    <canvas bind:this={canvasEl} class="liquid-canvas"></canvas>

    <!-- Overlay Stats & Target Details -->
    <div class="liquid-overlay">
      <div class="target-header">
        <Trophy size={20} class="trophy-icon" />
        <span class="target-title">{i18n.t.stats.liquidTargetTitle}</span>
      </div>

      <div class="cash-accumulated">
        <span class="currency">€</span>
        <span class="amount">
          {(appState.overallStats?.totalMoneyEarned || 0).toLocaleString(i18n.currentLocale === 'ru' ? 'ru-RU' : i18n.currentLocale === 'en' ? 'en-US' : 'it-IT', {
            minimumFractionDigits: 2,
            maximumFractionDigits: 2
          })}
        </span>
      </div>

      <div class="target-badge">
        <Sparkles size={14} />
        <span>
          Target: € {(appState.overallStats?.revenueTarget || 20000).toLocaleString(i18n.currentLocale === 'ru' ? 'ru-RU' : i18n.currentLocale === 'en' ? 'en-US' : 'it-IT')}
        </span>
      </div>

      <div class="percentage-pill">
        <span class="pct-val">
          {(appState.overallStats?.liquidFillPercent || 0).toFixed(1)}%
        </span>
        <span class="pct-label">{i18n.t.stats.liquidProgressBadge}</span>
      </div>
    </div>
  </div>
</div>

<style>
  .liquid-wrapper {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 30px;
  }

  .liquid-container {
    position: relative;
    width: 100%;
    max-width: 440px;
    height: 520px;
    border-radius: 28px;
    background: radial-gradient(circle at 50% 10%, rgba(21, 26, 53, 0.7) 0%, rgba(8, 10, 20, 0.95) 100%);
    border: 2px solid rgba(147, 168, 255, 0.2);
    box-shadow:
      0 20px 50px rgba(0, 0, 0, 0.7),
      inset 0 0 30px rgba(51, 82, 255, 0.15),
      0 0 30px rgba(51, 82, 255, 0.2);
    overflow: hidden;
  }

  .liquid-canvas {
    width: 100%;
    height: 100%;
    display: block;
  }

  .milestones-bar {
    position: absolute;
    right: 14px;
    top: 0;
    bottom: 0;
    width: 45px;
    pointer-events: none;
    z-index: 5;
  }

  .milestone-mark {
    position: absolute;
    right: 0;
    display: flex;
    align-items: center;
    gap: 5px;
    font-size: 10px;
    font-weight: 600;
    color: rgba(147, 168, 255, 0.45);
  }

  .milestone-mark::after {
    content: '';
    width: 8px;
    height: 1.5px;
    background: rgba(147, 168, 255, 0.3);
  }

  .liquid-overlay {
    position: absolute;
    inset: 0;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    pointer-events: none;
    z-index: 10;
    padding: 24px;
    text-align: center;
    gap: 12px;
  }

  .target-header {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 6px 14px;
    border-radius: 20px;
    background: rgba(8, 10, 20, 0.65);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    border: 1px solid rgba(147, 168, 255, 0.2);
  }

  :global(.trophy-icon) {
    color: #fbbf24;
    filter: drop-shadow(0 0 8px rgba(251, 191, 36, 0.6));
  }

  .target-title {
    font-size: 12px;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.6px;
    color: var(--accent-highlight);
  }

  .cash-accumulated {
    display: flex;
    align-items: baseline;
    gap: 6px;
    text-shadow: 0 4px 18px rgba(0, 0, 0, 0.8), 0 0 24px rgba(51, 82, 255, 0.5);
  }

  .currency {
    font-size: 26px;
    font-weight: 700;
    color: var(--accent-highlight);
  }

  .amount {
    font-size: 44px;
    font-weight: 800;
    letter-spacing: -1px;
    color: #ffffff;
  }

  .target-badge {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    font-weight: 600;
    color: var(--text-secondary);
    padding: 5px 14px;
    background: rgba(14, 18, 36, 0.6);
    backdrop-filter: blur(8px);
    -webkit-backdrop-filter: blur(8px);
    border-radius: 8px;
    border: 1px solid rgba(147, 168, 255, 0.15);
  }

  .percentage-pill {
    margin-top: 10px;
    display: flex;
    align-items: baseline;
    gap: 6px;
    padding: 8px 18px;
    background: rgba(8, 10, 20, 0.75);
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
    border-radius: 30px;
    border: 1px solid var(--accent-primary);
    box-shadow: 0 0 20px rgba(51, 82, 255, 0.4);
  }

  .pct-val {
    font-size: 22px;
    font-weight: 800;
    color: #ffffff;
  }

  .pct-label {
    font-size: 12px;
    font-weight: 600;
    color: var(--accent-highlight);
    text-transform: uppercase;
  }
</style>
