<script lang="ts">
    let { opacity = 0.7 } = $props();
</script>

<div class="absolute inset-0 overflow-hidden pointer-events-none -z-2">
<svg
    class="w-full h-full"
    viewBox="0 0 1600 900"
    preserveAspectRatio="xMidYMid slice"
    style:opacity={opacity}
>

<defs>

    <!-- Large Glow -->
    <filter id="glow-large" x="-100%" y="-100%" width="300%" height="300%">
        <feGaussianBlur stdDeviation="10"/>
    </filter>

    <!-- Medium Glow -->
    <filter id="glow-medium" x="-100%" y="-100%" width="300%" height="300%">
        <feGaussianBlur stdDeviation="30"/>
    </filter>

    <!-- Small Glow -->
    <filter id="glow-small" x="-100%" y="-100%" width="300%" height="300%">
        <feGaussianBlur stdDeviation="5"/>
    </filter>

    <!-- Cyan -->
    <linearGradient id="cyan">
        <stop offset="0%" stop-color="#00ADD8"/>
        <stop offset="35%" stop-color="#55E8FF"/>
        <stop offset="50%" stop-color="#FFFFFF"/>
        <stop offset="65%" stop-color="#55E8FF"/>
        <stop offset="100%" stop-color="#00ADD8"/>
    </linearGradient>

    <!-- Pink -->
    <linearGradient id="pink">
        <stop offset="0%" stop-color="#FF3A8D"/>
        <stop offset="35%" stop-color="#FF78C7"/>
        <stop offset="50%" stop-color="#FFFFFF"/>
        <stop offset="65%" stop-color="#FF78C7"/>
        <stop offset="100%" stop-color="#FF3A8D"/>
    </linearGradient>

    <!-- Yellow -->
    <linearGradient id="yellow">
        <stop offset="0%" stop-color="#FFD23F"/>
        <stop offset="40%" stop-color="#FFE979"/>
        <stop offset="50%" stop-color="#FFFFFF"/>
        <stop offset="60%" stop-color="#FFE979"/>
        <stop offset="100%" stop-color="#FFD23F"/>
    </linearGradient>

</defs>

<!-- ====================================================== -->
<!-- CYAN -->
<!-- ====================================================== -->

<g>

    <!-- outer glow -->
    <path
        d="M-120 760 C250 560 520 760 900 430 C1200 160 1500 350 1780 120"
        stroke="url(#cyan)"
        stroke-width="80"
        opacity=".05"
        fill="none"
        filter="url(#glow-large)"
        stroke-linecap="round"
    />

    <!-- mid glow -->
    <path
        d="M-120 760 C250 560 520 760 900 430 C1200 160 1500 350 1780 120"
        stroke="url(#cyan)"
        stroke-width="42"
        opacity=".14"
        fill="none"
        filter="url(#glow-medium)"
        stroke-linecap="round"
    />

    <!-- inner glow -->
    <path
        d="M-120 760 C250 560 520 760 900 430 C1200 160 1500 350 1780 120"
        stroke="url(#cyan)"
        stroke-width="18"
        opacity=".65"
        fill="none"
        filter="url(#glow-small)"
        stroke-linecap="round"
    />

    <!-- white beam -->
    <path
        d="M-120 760 C250 560 520 760 900 430 C1200 160 1500 350 1780 120"
        stroke="white"
        stroke-opacity=".9"
        stroke-width="4"
        fill="none"
        stroke-linecap="round"
    />

    <!-- color line -->
    <path
        d="M-120 760 C250 560 520 760 900 430 C1200 160 1500 350 1780 120"
        stroke="url(#cyan)"
        stroke-width="2"
        fill="none"
        stroke-linecap="round"
    />

</g>

<!-- ====================================================== -->
<!-- PINK -->
<!-- ====================================================== -->

<g opacity=".95">

    <path
        d="M-180 260 C180 40 480 460 820 260 C1180 70 1470 270 1750 180"
        stroke="url(#pink)"
        stroke-width="75"
        opacity=".05"
        fill="none"
        filter="url(#glow-large)"
        stroke-linecap="round"
    />

    <path
        d="M-180 260 C180 40 480 460 820 260 C1180 70 1470 270 1750 180"
        stroke="url(#pink)"
        stroke-width="38"
        opacity=".14"
        fill="none"
        filter="url(#glow-medium)"
        stroke-linecap="round"
    />

    <path
        d="M-180 260 C180 40 480 460 820 260 C1180 70 1470 270 1750 180"
        stroke="url(#pink)"
        stroke-width="16"
        opacity=".65"
        fill="none"
        filter="url(#glow-small)"
        stroke-linecap="round"
    />

    <path
        d="M-180 260 C180 40 480 460 820 260 C1180 70 1470 270 1750 180"
        stroke="white"
        stroke-opacity=".9"
        stroke-width="4"
        fill="none"
        stroke-linecap="round"
    />

    <path
        d="M-180 260 C180 40 480 460 820 260 C1180 70 1470 270 1750 180"
        stroke="url(#pink)"
        stroke-width="2"
        fill="none"
        stroke-linecap="round"
    />

</g>

<!-- ====================================================== -->
<!-- YELLOW -->
<!-- ====================================================== -->

<g>

    <path
        d="M1750 820 C1200 700 900 900 640 620 C420 430 180 650 -120 480"
        stroke="url(#yellow)"
        stroke-width="75"
        opacity=".04"
        fill="none"
        filter="url(#glow-large)"
        stroke-linecap="round"
    />

    <path
        d="M1750 820 C1200 700 900 900 640 620 C420 430 180 650 -120 480"
        stroke="url(#yellow)"
        stroke-width="40"
        opacity=".13"
        fill="none"
        filter="url(#glow-medium)"
        stroke-linecap="round"
    />

    <path
        d="M1750 820 C1200 700 900 900 640 620 C420 430 180 650 -120 480"
        stroke="url(#yellow)"
        stroke-width="18"
        opacity=".65"
        fill="none"
        filter="url(#glow-small)"
        stroke-linecap="round"
    />

    <path
        d="M1750 820 C1200 700 900 900 640 620 C420 430 180 650 -120 480"
        stroke="white"
        stroke-opacity=".9"
        stroke-width="4"
        fill="none"
        stroke-linecap="round"
    />

    <path
        d="M1750 820 C1200 700 900 900 640 620 C420 430 180 650 -120 480"
        stroke="url(#yellow)"
        stroke-width="2"
        fill="none"
        stroke-linecap="round"
    />

</g>

</svg>
</div>