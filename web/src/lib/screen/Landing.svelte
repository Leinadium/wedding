<script lang="ts">
  import l1 from "../../assets/landing/l1.png";
  import l2 from "../../assets/landing/l2.png";
  import l3 from "../../assets/landing/l3.png";
  import l4 from "../../assets/landing/l4.png";
  import l5 from "../../assets/landing/l5.png";
  import l6 from "../../assets/landing/l6.png";

  import { fly } from "svelte/transition";
  import Invite from "../guests/Invite.svelte";

  let invite = $state(false);

  function toggleInvite() {
    invite = !invite;
  }
</script>

<div id="landing">
  <img
    id="l1"
    src={l1}
    alt="card"
    transition:fly={{ delay: 150, y: 100, duration: 1000 }}
  />
  <img
    id="l2"
    src={l2}
    alt="invite"
    transition:fly={{ delay: 250, y: 120, duration: 1100 }}
  />
  <img
    id="l3"
    src={l3}
    alt="couple"
    transition:fly={{ delay: 500, y: 80, duration: 500 }}
  />
  <button id="l4" onclick={toggleInvite}>
    <img
      class="inside"
      src={l4}
      alt="details"
      transition:fly={{ delay: 750, y: 120, duration: 800 }}
    />
  </button>
  <img
    id="l5"
    src={l5}
    alt="couple2"
    transition:fly={{ delay: 850, y: 50, duration: 700 }}
  />
  <img
    id="l6"
    src={l6}
    alt="stamp"
    transition:fly={{ delay: 900, y: 60, duration: 900 }}
  />

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
    right: 2%;
    top: 15%;
    width: 60%;
    z-index: 2;
  }

  #l3 {
    left: 2%;
    top: 34%;
    width: 30%;
    transform: rotate(10deg);
    z-index: 3;
  }

  #l4 {
    left: 2%;
    width: 55%;
    top: 49%;
    z-index: 4;
  }

  #l5 {
    right: 7%;
    width: 34%;
    top: 52%;
    transform: rotate(-15deg);
    z-index: 5;
  }

  #l6 {
    right: 20%;
    width: 25%;
    top: 66%;
    transform: rotate(-15deg);
    z-index: 6;
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
