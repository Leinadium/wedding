<script lang="ts">
  import { fly } from "svelte/transition";

  let {
    src,
    alt,
    delay,
    y,
    duration,
  }: {
    src: string;
    alt: string;
    delay: number;
    y: number;
    duration: number;
  } = $props();

  let loaded = $state(false);
</script>

<img {src} alt="" onload={() => (loaded = true)} style="display: none;" />

{#if loaded}
  <img
    {src}
    {alt}
    transition:fly={{ delay: delay, y: y, duration: duration }}
  />
{/if}

<style>
  img {
    width: 100%;
  }
</style>
