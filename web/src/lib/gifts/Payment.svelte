<script lang="ts">
  import { onMount } from "svelte";
  import { api, FRONT_URL, type Purchase } from "../api";

  let {
    url,
    callback,
  }: {
    url: string | null;
    callback: (data: Purchase | null) => void;
  } = $props();

  let popup = $state<WindowProxy | null>();

  async function handleToken(token: string) {
    console.log(token);
    let purchase: Purchase | null = null;
    try {
      const res = await api.getPurchase(token);
      purchase = res.purchase;
    } catch (e) {
      console.error(e);
    }
    callback(purchase);
  }

  function handleMessage(event: MessageEvent) {
    if (event.origin != FRONT_URL) return;
    if (event.data.callback) {
      handleToken(event.data.callback);
    }
    popup!.close();
  }

  onMount(() => {
    if (url) {
      popup = window.open(url, "popup", "popup=true");
    }
  });
</script>

<svelte:window onmessage={handleMessage} />
