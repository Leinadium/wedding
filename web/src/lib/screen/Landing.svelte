<script lang="ts">
  import l1 from "../../assets/landing/l1.png";
  import l2 from "../../assets/landing/l2.png";
  import l3 from "../../assets/landing/l3.png";
  import l4 from "../../assets/landing/l4.png";
  import l5 from "../../assets/landing/l5.png";
  import l6 from "../../assets/landing/l6.png";
  import l7 from "../../assets/landing/l7.png";
  import l8 from "../../assets/landing/l8.png";

  import { fly } from "svelte/transition";
  import Invite from "../guests/Invite.svelte";
    import Image from "./Image.svelte";

  let invite = $state(false);

  function toggleInvite() {
    invite = !invite;
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
        <Image src={l3} alt="couple" delay={500} y={80} duration={500} />
    </div>

    <button id="l4" onclick={toggleInvite}>
        <div class="inside">
            <Image src={l4} alt="details" delay={750} y={120} duration={800} />
        </div>
    </button>

    <div id="l5" class="image-container">
        <Image src={l5} alt="couple2" delay={850} y={50} duration={700} />
    </div>

    <div id="l6" class="image-container">
        <Image src={l6} alt="stamp" delay={1000} y={60} duration={600} />
    </div>

    <div id="l7" class="image-container">
        <Image src={l7} alt="photo3" delay={1100} y={60} duration={600} />
    </div>

    <div id="l8" class="image-container">
        <Image src={l8} alt="portal" delay={1250} y={30} duration={600} />
    </div>

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
    top: 27%;
    width: 30%;
    transform: rotate(10deg);
    z-index: 3;
  }

  #l4 {
    left: 0%;
    width: 55%;
    top: 39%;
    z-index: 4;
  }

  #l5 {
    right: 7%;
    width: 34%;
    top: 43%;
    transform: rotate(-15deg);
    z-index: 5;
  }

  #l6 {
    right: 20%;
    width: 25%;
    top: 54%;
    transform: rotate(-15deg);
    z-index: 6;
  }
  #l7 {
    left: 3%;
    width: 29%;
    top: 56%;
    transform: rotate(7deg);
    z-index: 7;
  }
  #l8 {
    right: 3%;
    width: 50%;
    top: 60%;
    z-index: 7;
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

  button {
    background: none;
    border: none;
    padding: 0;
    cursor: pointer;
    position: absolute;
  }
</style>
