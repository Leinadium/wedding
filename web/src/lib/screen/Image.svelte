<script lang="ts">
  import { fly } from "svelte/transition";

  let {
    src,
    alt,
    delay,
    y,
    duration,
  }: {
    src: string,
    alt: string,
    delay: number,
    y: number,
    duration: number,
  } = $props();

  let loaded = $state(false)
</script>


<img
    src={src}
    alt=""
    onload={() => loaded = true}
    style="display: none;"
/>

{#if loaded}
	<img
		src={src}
		alt={alt}
		transition:fly={{ delay: delay, y: y, duration: duration }}
	/>
{/if}

<style>
  img {
      width: 100%;
      filter: drop-shadow(0px 20px 40px rgba(94, 58, 73, 0.5));
      transition: filter 0.3s ease;
  }

  img:hover{
      filter: drop-shadow(0px 20px 30px rgba(94, 58, 73, 0.7));
  }


</style>
