<script lang="ts">
  import Header from "./lib/landing/Header.svelte";
  import Footer from "./lib/landing/Footer.svelte";
  import Landing from "./lib/landing/Landing.svelte";
  import Callback from "./lib/Callback.svelte";
  import { onMount } from "svelte";
  import { getText } from "./lib/text/text";
  import { Page } from "./lib/common";
  import Story from "./lib/Story.svelte";
  import Backoffice from "./lib/backoffice/Backoffice.svelte";

  let backoffice = $state(true);
  let showContent = $state(true);
  onMount(() => {
    const thisUrl = new URLSearchParams(window.location.search);
    showContent = !thisUrl.has("callback");
    backoffice = thisUrl.has("backoffice");
  });

  let page: Page = $state(Page.Landing);

  function setState(p: Page) {
    page = p;
  }
</script>

<svelte:head>
  <title>{getText("site-title")}</title>
</svelte:head>

{#if backoffice}
  <Backoffice />
{:else if showContent}
  <main>
    <Header />
    {#if page === Page.Landing}
      <Landing {setState} />
    {:else}
      <Story {setState} />
    {/if}
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
    font-family: "Bodoni Moda", Courier, monospace;
  }

  main {
    position: relative;
    min-height: 100dvh;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 2rem 1rem;
    padding-bottom: 0;
    overflow-x: hidden;
  }
</style>
