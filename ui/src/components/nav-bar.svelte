<script lang="ts">
    import { AppBar } from '@skeletonlabs/skeleton';
  import FontAwesomeIcon from './font-awesome-icon.svelte';
</script>
<AppBar background="bg-surface-900" padding="sm:px-10 px-5 py-3 md:px-20 sm:py-5" spacing="">
    <div slot="lead">
        <a href="/" class="text-white font-extrabold">
            DEVEMM
        </a>
    </div>
    <div class="hidden sm:inline-flex px-0x py-1 mx-10x justify-center items-center gap-12 w-full">
        <div class="nav-link">
            <a href="/about" class="text-white">
                Home
            </a>
        </div>
        <div class="nav-link">
            <a href="/about" class="text-white">
                Portfolio
            </a>
        </div>
        <!-- <div class="nav-link">
            <a href="/about" class="text-white">
                Courses
            </a>
        </div> -->
        <div class="nav-link">
            <a href="/about" class="text-white">
                Blog
            </a>
        </div>
        <div class="nav-link">
            <a href="/about" class="text-white">
                Contact
            </a>
        </div>
    </div>
    <div slot="trail" class="flex justify-between items-center">
        <div class="mr-2">
            <button class="btn variant-filled bg-primary-500 max-md:btn-sm">Hire me</button>
        </div>
        <div class="ml-2">
            <FontAwesomeIcon icon="fa-solid fa-bars " className="text-white text-2xl sm:hidden" />
        </div>
    </div>
</AppBar>