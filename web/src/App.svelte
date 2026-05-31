<script lang="ts">
  import Invite from "./lib/guests/Invite.svelte";
  import Gifts from "./lib/gifts/Gifts.svelte";
  import Header from "./lib/screen/Header.svelte";
  import Footer from "./lib/screen/Footer.svelte";
  // 1. Logic for collapsible sections
  let rsvpOpen = $state(false);
  let giftsOpen = $state(false);
  let locationOpen = $state(false);

  function toggleRSVP() {
    rsvpOpen = !rsvpOpen;
  }
  function toggleGifts() {
    giftsOpen = !giftsOpen;
  }
  function toggleLocation() {
    locationOpen = !locationOpen;
  }
</script>

<main>
  <Header />

  <div class="content">
    <h1 class="title">Our Big Day</h1>

    <section class="intro">
      <p>We're so excited to celebrate with you!</p>
      <p>Join us on Month 00th, 2027.</p>
    </section>

    <!-- RSVP wrapper -->
    <div class="collapsible-wrapper">
      <button onclick={toggleRSVP} class="toggle-btn" aria-expanded={rsvpOpen}>
        RSVP {rsvpOpen ? "-" : "+"}
      </button>

      {#if rsvpOpen}
        <Invite />
      {/if}
    </div>

    <!-- Gifts wrapper -->
    <div class="collapsible-wrapper">
      <button
        onclick={toggleGifts}
        class="toggle-btn"
        aria-expanded={giftsOpen}
      >
        Gifts {giftsOpen ? "-" : "+"}
      </button>

      {#if giftsOpen}
        <Gifts />
      {/if}
    </div>
  </div>

  <Footer />
</main>

<style>
  :global(body) {
    margin: 0;
    background-color: #fffaf5;
    color: #4a4a4a;
    font-family: "Georgia", serif;
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

  .content {
    z-index: 1;
    width: 100%;
    max-width: 700px;
    text-align: center;
  }

  .title {
    font-size: 2.5rem;
    font-weight: normal;
    margin-bottom: 1.5rem;
  }

  .collapsible-wrapper {
    margin: 1.5rem 0;
    border-top: 1px solid #eee;
  }

  .toggle-btn {
    width: 100%;
    background: none;
    border: none;
    padding: 1.5rem 0;
    font-size: 1.2rem;
    font-family: inherit;
    cursor: pointer;
    color: #8a7b6e;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
</style>
