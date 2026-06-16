<script lang="tsx">
import { defineComponent, computed } from 'vue'
import { Collapse } from '@/components/Collapse'
import { LocaleDropdown } from '@/components/LocaleDropdown'
import { SizeDropdown } from '@/components/SizeDropdown'
import { UserInfo } from '@/components/UserInfo'
import { Screenfull } from '@/components/Screenfull'
import { Breadcrumb } from '@/components/Breadcrumb'
import { NotifyDropdown } from '@/components/NotifyDropdown'
import { useAppStore } from '@/stores/app'
import { useDesign } from '@/hooks/web/useDesign'
import Icon from '@/components/Icon/src/Icon.vue'
import { useTagsViewStore } from '@/stores/tagsView'
import { useZoneStoreWithOut } from '@/stores/zone'
import { useNow } from '@/hooks/web/useNow'
import { getRole } from '@/utils/roleUtils'

const { showTime } = useNow(true)
export default defineComponent({
  name: 'ToolHeader',
  setup() {
    const { getPrefixCls, variables } = useDesign()

    const prefixCls = getPrefixCls('tool-header')

    const appStore = useAppStore()
    const zoneStore = useZoneStoreWithOut()
    
    const breadcrumb = computed(() => appStore.getBreadcrumb)

    
    const hamburger = computed(() => appStore.getHamburger)

    
    const screenfull = computed(() => appStore.getScreenfull)

    
    const size = computed(() => appStore.getSize)

    
    const layout = computed(() => appStore.getLayout)

    
    const locale = computed(() => appStore.getLocale)

    const storeClear = async () => {
      const tagsViewStore = useTagsViewStore()
      tagsViewStore.delAllViews()
      const user: any = localStorage.getItem('user')
      const public_Key: any = localStorage.getItem('public_Key')
      const lang: any = localStorage.getItem('lang')
      const cdn: any = localStorage.getItem('cdn')
      const refresh_cdn: any = localStorage.getItem('refresh_cdn')
      const notify_message: any = localStorage.getItem('notify_message')
      localStorage.clear()
      localStorage.setItem('user', user)
      localStorage.setItem('public_Key', public_Key)
      localStorage.setItem('lang', lang)
      localStorage.setItem('cdn', cdn)
      localStorage.setItem('refresh_cdn', refresh_cdn)
      localStorage.setItem('notify_message', notify_message)
      await getRole()
      window.location.reload()
    }
    return {
      prefixCls,
      variables,
      breadcrumb,
      hamburger,
      screenfull,
      size,
      layout,
      locale,
      showTime,
      storeClear,
      zoneStore
    }
  },
  render() {
    return (
      <div
        id={`${this.variables.namespace}-tool-header`}
        class={[
          this.prefixCls,
          'h-[var(--top-tool-height)] relative px-[var(--top-tool-p-x)] flex items-center justify-between'
        ]}
      >
        {this.layout !== 'top' ? (
          <div class="h-full flex items-center">
            {this.hamburger && this.layout !== 'cutMenu' ? (
              <Collapse class="custom-hover" color="var(--top-header-text-color)"></Collapse>
            ) : undefined}
            {this.breadcrumb ? <Breadcrumb class="<md:hidden"></Breadcrumb> : undefined}
          </div>
        ) : undefined}
        <div class="h-full flex items-center">
          <div class="custom-hover">
            <span>
              (UTC{this.zoneStore.getZone && this.zoneStore.getZone.indexOf('-') == -1 ? '+' : ''}{' '}
              {this.zoneStore.getZone}){' '}
            </span>
            <span class="ml-2">{showTime.value}</span>
          </div>
          <div class="custom-hover" onClick={this.storeClear}>
            <Icon icon="material-symbols:database-off-rounded" class="mr-1" size={18} />
          </div>
          {this.screenfull ? (
            <Screenfull class="custom-hover" color="var(--top-header-text-color)"></Screenfull>
          ) : undefined}
          <NotifyDropdown class="custom-hover" />
          {this.size ? (
            <SizeDropdown class="custom-hover" color="var(--top-header-text-color)"></SizeDropdown>
          ) : undefined}
          {this.locale ? (
            <LocaleDropdown
              class="custom-hover"
              color="var(--top-header-text-color)"
            ></LocaleDropdown>
          ) : undefined}
          <UserInfo></UserInfo>
        </div>
      </div>
    )
  }
})
</script>

<style lang="less" scoped>
@prefix-cls: ~'@{adminNamespace}-tool-header';

.@{prefix-cls} {
  transition: left var(--transition-time-02);
}
</style>

