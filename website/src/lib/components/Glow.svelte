<script lang="ts">
    let {
        opacity = 0.55,
        dark = "dark"
    } = $props();

    const ribbon = dark === "dark"
        ? "#22c7ff"
        : "#0ea5e9";
</script>

<div class="fixed inset-0 pointer-events-none z-10 overflow-hidden">
<svg
    class="w-full h-full"
    viewBox="0 0 1600 900"
    preserveAspectRatio="xMidYMid slice"
    style:opacity={opacity}
>

<defs>

    <!-- Very soft blur -->
    <filter
        id="blur-xl"
        x="-50%"
        y="-50%"
        width="200%"
        height="200%"
    >
        <feGaussianBlur stdDeviation="70"/>
    </filter>

    <!-- Medium blur -->
    <filter
        id="blur-lg"
        x="-50%"
        y="-50%"
        width="200%"
        height="200%"
    >
        <feGaussianBlur stdDeviation="30"/>
    </filter>

    <!-- Slight blur -->
    <filter
        id="blur-sm"
        x="-50%"
        y="-50%"
        width="200%"
        height="200%"
    >
        <feGaussianBlur stdDeviation="8"/>
    </filter>

</defs>


<!-- ========================================= -->
<!-- Secondary Ribbon -->
<!-- ========================================= -->

<!-- Huge -->
<path
    d="
    M-250 760
    C120 640
    360 980
    760 860
    "
    fill="none"
    stroke={ribbon}
    stroke-width="140"
    opacity=".025"
    filter="url(#blur-xl)"
    stroke-linecap="round"
/>

<!-- Medium -->
<path
    d="
    M-250 760
    C120 640
    360 980
    760 860
    "
    fill="none"
    stroke={ribbon}
    stroke-width="70"
    opacity=".06"
    filter="url(#blur-lg)"
    stroke-linecap="round"
/>

<!-- Inner -->
<path
    d="
    M-250 760
    C120 640
    360 980
    760 860
    "
    fill="none"
    stroke={ribbon}
    stroke-width="30"
    opacity=".12"
    filter="url(#blur-sm)"
    stroke-linecap="round"
/>

<!-- Core -->
<path
    d="
    M-250 760
    C120 640
    360 980
    760 860
    "
    fill="none"
    stroke={ribbon}
    stroke-width="7"
    opacity=".25"
    stroke-linecap="round"
/>


<!-- Huge dispersion -->
<path
    d="
    M760 980
    C980 760
    1260 430
    1820 140
    "
    fill="none"
    stroke={ribbon}
    stroke-width="180"
    opacity=".035"
    filter="url(#blur-xl)"
    stroke-linecap="round"
/>

<!-- Mid -->
<path
    d="
    M760 980
    C980 760
    1260 430
    1820 140
    "
    fill="none"
    stroke={ribbon}
    stroke-width="90"
    opacity=".08"
    filter="url(#blur-lg)"
    stroke-linecap="round"
/>

<!-- Inner -->
<path
    d="
    M760 980
    C980 760
    1260 430
    1820 140
    "
    fill="none"
    stroke={ribbon}
    stroke-width="42"
    opacity=".14"
    filter="url(#blur-sm)"
    stroke-linecap="round"
/>

<!-- Core -->
<path
    d="
    M760 980
    C980 760
    1260 430
    1820 140
    "
    fill="none"
    stroke={ribbon}
    stroke-width="10"
    opacity=".35"
    stroke-linecap="round"
/>

</svg>
</div>