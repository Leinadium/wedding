<script lang="ts">
  import l1 from "../../assets/landing/l1.png";
  import l2 from "../../assets/landing/l2.png";
  import l3 from "../../assets/landing/l3.png";
  import l4 from "../../assets/landing/l4.png";
  import l5 from "../../assets/landing/l5.png";
  import l6 from "../../assets/landing/l6.png";
  import l7 from "../../assets/landing/l7.png";
  import l8 from "../../assets/landing/l8.png";
  import l9 from "../../assets/landing/l9.png";

  import Invite from "../guests/Invite.svelte";
  import Image from "./Image.svelte";
  import Gifts from "../gifts/Gifts.svelte";
  import Venue from "../venue/Venue.svelte";
  import { getTextDefault } from "../text/text";

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
  <!-- <button id="l5" onclick={toggleGifts}>
    <div class="inside">
      <Image src={l5} alt="portal" delay={820} y={30} duration={600} />
    </div>
  </button> -->

  <div id="l6" class="image-container">
    <Image src={l6} alt="stamp" delay={1000} y={60} duration={600} />
  </div>

  <div id="l7" class="image-container">
    <Image src={l7} alt="photo3" delay={1100} y={60} duration={600} />
  </div>

  <div id="l8" class="image-container">
    <Image src={l8} alt="couple2" delay={1200} y={50} duration={700} />
  </div>

  <button id="l9" onclick={toggleVenue}>
    <div class="inside">
      <Image src={l9} alt="letter" delay={1350} y={50} duration={600} />
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
    aspect-ratio: 1 / 3;
  }

  #l1 {
    left: 0%;
    top: 0;
    width: 60%;
    z-index: 1;
  }

  #l2 {
    right: 2%;
    top: 13%;
    width: 60%;
    z-index: 2;
  }

  #l3 {
    left: 2%;
    top: 25%;
    width: 30%;
    transform: rotate(10deg);
    z-index: 3;
  }

  #l4 {
    left: 0%;
    width: 55%;
    top: 35%;
    z-index: 4;
  }

  #l5 {
    right: 3%;
    width: 50%;
    top: 40%;
    transform: rotate(-3deg);
    z-index: 5;
  }

  #l6 {
    left: 3%;
    width: 34%;
    top: 50%;
    transform: rotate(7deg);
    z-index: 6;
  }

  #l7 {
    left: 30%;
    width: 25%;
    top: 58%;
    transform: rotate(-15deg);
    z-index: 10;
  }

  #l8 {
    right: 7%;
    width: 37%;
    top: 65%;
    transform: rotate(15deg);
    z-index: 10;
  }

  #l9 {
    left: 0%;
    width: 60%;
    top: 60%;
    transform: rotate(1deg);
    z-index: 9;
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
