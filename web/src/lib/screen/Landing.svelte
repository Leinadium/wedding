<script lang="ts">
  import { fly } from "svelte/transition";
  import Invite from "../guests/Invite.svelte";
  import { cubicIn } from "svelte/easing";

  let invite = $state(false);

  function toggleInvite() {
    invite = !invite;
  }
</script>

<div id="landing">
  <img
    id="l1"
    src="src/assets/landing/l1.png"
    alt="card"
    transition:fly={{ delay: 150, y: 100, duration: 1000 }}
  />
  <img
    id="l2"
    src="src/assets/landing/l2.png"
    alt="invite"
    transition:fly={{ delay: 250, y: 120, duration: 1100 }}
  />
  <img
    id="l3"
    src="src/assets/landing/l3.png"
    alt="couple"
    transition:fly={{ delay: 500, y: 80, duration: 500 }}
  />
  <button id="l4" onclick={toggleInvite}>
    <img
      class="inside"
      src="src/assets/landing/l4.png"
      alt="details"
      transition:fly={{ delay: 750, y: 120, duration: 800 }}
    />
  </button>

  {#if invite}
    <Invite closeCb={toggleInvite} />
  {/if}
</div>

<style>
  #landing {
    position: relative;

    width: 100%;
    max-width: 800px;
    height: 100%;
    aspect-ratio: 1 / 2.5;
  }

  #l1 {
    left: 0%;
    top: 0;
    width: 60%;
    z-index: 1;
  }

  #l2 {
    right: 0%;
    top: 10%;
    width: 60%;
    z-index: 2;
  }

  #l3 {
    left: 2%;
    top: 31%;
    width: 30%;
    transform: rotate(10deg);
    z-index: 3;
  }

  #l4 {
    left: 3%;
    width: 50%;
    top: 44%;
    z-index: 4;
  }

  #landing > img {
    position: absolute;
    height: auto;
  }

  .inside {
    width: 100%;
    height: auto;
    display: block;
  }

  button {
    background: none;
    border: none;
    padding: 0;
    cursor: pointer;
    position: absolute;
  }
</style>
