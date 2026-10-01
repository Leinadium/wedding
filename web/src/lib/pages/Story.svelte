<script lang="ts">
  import { fade } from "svelte/transition";

  import headerImg from "../../assets/story/header.jpg";
  import p0 from "../../assets/story/p0.png";
  import p1 from "../../assets/story/p1.png";
  import p2 from "../../assets/story/p2.png";
  import p3 from "../../assets/story/p3.png";
  import p4 from "../../assets/story/p4.png";
  import p5 from "../../assets/story/p5.png";

  import { Page } from "../common";
  import Image from "../Image.svelte";
  import Text from "../text/Text.svelte";
  import { onDestroy, onMount } from "svelte";
  import BackButton from "./BackButton.svelte";
  import Header from "./Header.svelte";
  import { getTextDefault } from "../text/text";
  let {
    setState,
  }: {
    setState: (page: Page) => void;
  } = $props();

  function back() {
    setState(Page.Landing);
  }

  function scrollTop() {
    window.scrollTo({ top: 0, behavior: "smooth" });
  }

  onMount(scrollTop);
  onDestroy(scrollTop);
</script>

<div class="story">
  <Header
    src={headerImg}
    text={getTextDefault("story-title", "Our Story")}
    offset="8%"
  />

  <p class="formal first" in:fade={{ delay: 400, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text01" /><br />
    &emsp;&emsp;<Text key="story-text02" /><br />
    &emsp;&emsp;<Text key="story-text03" /><br />
    &emsp;&emsp;<Text key="story-text04" /><br />
    &emsp;&emsp;<Text key="story-text05" /><br />
  </p>

  <div id="p1" class="photo" style="transform: rotate(5deg);">
    <Image src={p1} alt="p1" delay={600} y={100} duration={1000} />
  </div>

  <p class="formal" in:fade={{ delay: 800, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text06" /><br />
    &emsp;&emsp;<Text key="story-text07" /><br />
  </p>

  <div id="p2" class="photo" style="transform: rotate(-5deg);">
    <Image src={p2} alt="p2" delay={1000} y={100} duration={1000} />
  </div>

  <p class="formal" in:fade={{ delay: 1200, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text08" /><br />
    &emsp;&emsp;<Text key="story-text09" /><br />
    &emsp;&emsp;<Text key="story-text10" /><br />
  </p>

  <div id="p3" class="photo" style="transform: rotate(5deg);">
    <Image src={p3} alt="p3" delay={1400} y={100} duration={1000} />
  </div>

  <p class="formal" in:fade={{ delay: 1600, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text11" /><br />
    &emsp;&emsp;<Text key="story-text12" /><br />
  </p>

  <div id="p4" class="photo" style="transform: rotate(-5deg);">
    <Image src={p4} alt="p4" delay={1800} y={100} duration={1000} />
  </div>

  <p class="formal quote" in:fade={{ delay: 2000, duration: 200 }}>
    &emsp;&emsp;<Text key="story-text13" /><br />
    &emsp;&emsp;&emsp;&emsp; - <Text key="story-text14" /><br />
    &emsp;&emsp;<Text key="story-text15" /><br />
    &emsp;&emsp;&emsp;&emsp; - <Text key="story-text16" /><br />
  </p>

  <div id="p5" class="photo" style="transform: rotate(-5deg);">
    <Image src={p5} alt="p5" delay={2200} y={100} duration={1000} />
  </div>

  <BackButton textKey="story-back" action={back} color="#f9f6f1" />
</div>

<style>
  .story {
    display: flex;
    flex-flow: column nowrap;
    align-items: center;

    min-height: 90vh;
    width: 100vw;

    color: #f9f6f1;
    font-size: 1.2rem;
    text-align: justify;
  }

  p {
    width: 80%;
    max-width: 600px;
  }

  .first {
    margin-top: 1rem;
  }

  .photo {
    height: auto;
    width: 30vh;
  }

  .quote {
    font-style: italic;
  }
</style>
