<script lang="ts">
  import Header from "./lib/screen/Header.svelte";
  import Footer from "./lib/screen/Footer.svelte";
  import Countdown from "./lib/countdown/Countdown.svelte";
  import Landing from "./lib/screen/Landing.svelte";
  import Callback from "./lib/callback/Callback.svelte";
  import { onMount } from "svelte";

  let showContent = $state(true);
  onMount(() => {
    const thisUrl = new URLSearchParams(window.location.search);
    console.log(thisUrl);
    showContent = !thisUrl.has("callback");
  });
</script>

{#if showContent}
  <main>
    <Header />
    <Landing />
    <Countdown />
    <Footer />
  </main>
{:else}
  <Callback />
{/if}

<style>
  :global(body) {
    margin: 0;
    color: #4a4a4a;
    font-family: "Georgia", serif;

    background-color: #c79ead;
    background-image:
      url("data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='noiseFilter'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.75' numOctaves='3' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23noiseFilter)' opacity='0.08'/%3E%3C/svg%3E"),
      radial-gradient(circle at 50% 0%, #c79ead 0%, #a67b8b 100%);
  }

  :global(.cursive) {
    font-family: "Luxurious Script", cursive;
  }

  :global(.formal) {
    font-family: "Cormorant Garamond", Georgia, "Times New Roman", Times, serif;
  }
  :global(.formal-num) {
    font-family: "Courier New", Courier, monospace;
  }

  main {
    position: relative;
    min-height: 100dvh;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 2rem 1rem;
    overflow-x: hidden;
  }
</style>
