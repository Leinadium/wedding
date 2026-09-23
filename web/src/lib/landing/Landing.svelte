<script lang="ts">
  import l1 from "../../assets/landing/l1.png";
  import l1b from "../../assets/landing/l1b.png";
  import l2 from "../../assets/landing/l2.png";
  import l3 from "../../assets/landing/l3.png";
  import l3b from "../../assets/landing/l3b.png";
  import l4 from "../../assets/landing/l4.png";
  import l5 from "../../assets/landing/l5.png";
  import l6 from "../../assets/landing/l6.png";
  import l7 from "../../assets/landing/l7.png";
  import l8 from "../../assets/landing/l8.png";
  import l8b from "../../assets/landing/l8b.png";
  import l9 from "../../assets/landing/l9.png";
  import l10 from "../../assets/landing/l10.png";

  import Invite from "../guests/Invite.svelte";
  import Image from "../Image.svelte";
  import Gifts from "../gifts/Gifts.svelte";
  import Venue from "../Venue.svelte";
  import { getTextDefault } from "../text/text";
  import { Page } from "../common";

  let {
    setState,
  }: {
    setState: (page: Page) => void;
  } = $props();

  let invite = $state(false);
  let gifts = $state(false);
  let venue = $state(false);

  const giftsUrl: string = getTextDefault("gifts-url", "/#");

  function toggleInvite() {
    invite = !invite;
  }
  function toggleGifts() {
    gifts = !gifts;
  }
  function toggleVenue() {
    venue = !venue;
  }
</script>

<div id="landing">
  <div id="l1" class="image-container">
    <Image src={l1} alt="card" delay={150} y={100} duration={1000} />
  </div>

  <div id="l2" class="image-container">
    <Image src={l2} alt="invite" delay={250} y={120} duration={1100} />
  </div>

  <div id="l3" class="image-container">
    <Image src={l3} alt="couple" delay={500} y={80} duration={1200} />
  </div>

  <div id="l3b" class="image-container">
    <Image src={l3b} alt="couple-b" delay={700} y={80} duration={1200} />
  </div>

  <button id="l4" onclick={toggleInvite}>
    <div class="inside">
      <Image src={l4} alt="details" delay={750} y={120} duration={800} />
    </div>
  </button>

  <a id="l5" href={giftsUrl} target="_blank">
    <div class="inside">
      <Image src={l5} alt="portal" delay={820} y={30} duration={600} />
    </div>
  </a>

  <div id="l6" class="image-container">
    <Image src={l6} alt="stamp" delay={1000} y={60} duration={600} />
  </div>

  <div id="l7" class="image-container">
    <Image src={l7} alt="photo3" delay={1100} y={60} duration={600} />
  </div>

  <div id="l8" class="image-container">
    <Image src={l8} alt="couple2" delay={1200} y={50} duration={700} />
  </div>

  <div id="l8b" class="image-container">
    <Image src={l8b} alt="photo8b" delay={1400} y={60} duration={600} />
  </div>

  <button id="l9" onclick={toggleVenue}>
    <div class="inside">
      <Image src={l9} alt="letter" delay={1350} y={50} duration={600} />
    </div>
  </button>

  <button id="l10" onclick={() => setState(Page.Story)}>
    <div class="inside">
      <Image src={l10} alt="story" delay={1350} y={50} duration={600} />
    </div>
  </button>

  {#if invite}
    <Invite closeCb={toggleInvite} />
  {/if}

  {#if gifts}
    <Gifts closeCb={toggleGifts} />
  {/if}

  {#if venue}
    <Venue closeCb={toggleVenue} />
  {/if}
</div>

<style>
  #landing {
    position: relative;

    width: 100%;
    max-width: 800px;
    height: 100%;
    aspect-ratio: 1 / 3.2;
  }

  #l1 {
    left: 2%;
    top: 0;
    width: 60%;
    z-index: 1;
  }

  #l2 {
    right: -2%;
    top: 12%;
    width: 60%;
    z-index: 2;
  }

  #l3 {
    left: 0%;
    top: 23%;
    width: 30%;
    transform: rotate(10deg);
    z-index: 3;
  }

  #l3b {
    left: 14%;
    top: 30%;
    width: 30%;
    transform: rotate(-6deg);
    z-index: 3;
  }

  #l4 {
    left: 2%;
    width: 55%;
    top: 37%;
    transform: rotate(9deg);
    z-index: 4;
  }

  #l5 {
    right: 0%;
    width: 52%;
    top: 38%;
    transform: rotate(2deg);
    z-index: 5;
  }

  #l6 {
    left: 5%;
    width: 34%;
    top: 51%;
    transform: rotate(7deg);
    z-index: 6;
  }

  #l7 {
    left: 32%;
    width: 20%;
    top: 56%;
    transform: rotate(-15deg);
    z-index: 7;
  }

  #l8 {
    right: 5%;
    width: 33%;
    top: 60%;
    transform: rotate(15deg);
    z-index: 8;
  }

  #l8b {
    right: 5%;
    width: 33%;
    top: 69%;
    transform: rotate(11deg);
    z-index: 10;
  }

  #l9 {
    left: 2%;
    width: 60%;
    top: 60%;
    transform: rotate(1deg);
    z-index: 9;
  }

  #l10 {
    right: 5%;
    width: 50%;
    top: 78%;
    transform: rotate(7deg);
    z-index: 11;
  }

  #landing > .image-container {
    position: absolute;
    height: auto;
  }

  .inside {
    width: 100%;
    height: auto;
    display: block;
  }

  button,
  a {
    background: none;
    border: none;
    padding: 0;
    cursor: pointer;
    position: absolute;
  }
</style>
