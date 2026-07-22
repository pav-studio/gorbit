<script lang="ts">
	let {
		opacity = 0.45,
		dark = "dark"
	} = $props();

	let ribbon = $derived(
		dark === "dark" ? "#22c7ff" : "#0ea5e9"
	);
</script>

<div class="fixed inset-0 pointer-events-none overflow-hidden z-10">
	<svg
		class="w-full h-full"
		viewBox="0 0 1000 1000"
		preserveAspectRatio="none"
		style:opacity={opacity}
	>

		<defs>

			<filter
				id="glow"
				x="-20%"
				y="-20%"
				width="140%"
				height="140%"
			>
				<feGaussianBlur stdDeviation="18"/>
			</filter>

			<filter
				id="soft"
				x="-15%"
				y="-15%"
				width="130%"
				height="130%"
			>
				<feGaussianBlur stdDeviation="8"/>
			</filter>

		</defs>

		<!-- Glow -->

		<path
			d="
			M-120 820
			C180 640
			420 940
			710 720
			S1160 120
			"
			fill="none"
			stroke={ribbon}
			stroke-width="90"
			opacity=".035"
			filter="url(#glow)"
			stroke-linecap="round"
		/>

		<!-- Inner Glow -->

		<path
			d="
			M-120 820
			C180 640
			420 940
			710 720
			S1160 120
			"
			fill="none"
			stroke={ribbon}
			stroke-width="34"
			opacity=".12"
			filter="url(#soft)"
			stroke-linecap="round"
		/>

		<!-- Ribbon -->

		<path
			d="
			M-120 820
			C180 640
			420 940
			710 720
			S1160 120
			"
			fill="none"
			stroke={ribbon}
			stroke-width="8"
			opacity=".35"
			stroke-linecap="round"
		/>

		<!-- Secondary orbit -->

		<path
			d="
			M220 -80
			C420 120
			720 260
			1120 180
			"
			fill="none"
			stroke={ribbon}
			stroke-width="40"
			opacity=".03"
			filter="url(#soft)"
			stroke-linecap="round"
		/>

		<path
			d="
			M220 -80
			C420 120
			720 260
			1120 180
			"
			fill="none"
			stroke={ribbon}
			stroke-width="5"
			opacity=".18"
			stroke-linecap="round"
		/>

	</svg>
</div>