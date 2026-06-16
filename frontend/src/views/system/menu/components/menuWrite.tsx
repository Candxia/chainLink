import { h, reactive, ref } from 'vue'
import { Form, FormSchema } from '@/components/Form'
import { useValidator } from '@/hooks/web/useValidator'
import { useForm } from '@/hooks/web/useForm'
import { ElMessage } from 'element-plus'
import { DialogApi } from '@/components/MessageBox'
import { useI18n } from '@/hooks/web/useI18n'
import { fetchApiListApi, fetchMenuDetailApi, menuModifyApi } from '@/api/system'
import { IconPicker } from '@/components/IconPicker'
import { SelectHeader } from '@/components/Custom'

const { required } = useValidator()
const { t } = useI18n()

const colsSpan = {
  span: 24
}
const apiList = ref([])
const getApiList = async () => {
  const res = await fetchApiListApi()
  apiList.value = res.data.list
}
const menuList: any = ref([])

const schema = reactive<FormSchema[]>([
  {
    field: 'parent_id',
    label: t('system.menu.parentMenu'),
    colProps: colsSpan,
    component: 'TreeSelect',
    componentProps: {
      nodeKey: 'menu_id',
      props: {
        label: 'title',
        value: 'menu_id',
        children: 'children'
      },
      highlightCurrent: true,
      expandOnClickNode: false,
      checkStrictly: true,
      checkOnClickNode: true,
      disabled: true
    }
  },
  {
    field: 'menu_name',
    label: t('menu.menuName'),
    colProps: {
      span: 12
    },
    component: 'Input',
    formItemProps: {
      rules: [required()]
    }
  },
  {
    field: 'title',
    label: t('menu.menuName'),
    colProps: {
      span: 12
    },
    component: 'Input',
    formItemProps: {
      rules: [required()]
    }
  },
  {
    field: 'order_num',
    label: t('common.sort'),
    colProps: colsSpan,
    component: 'InputNumber',
    componentProps: {
      step: 1,
      min: 0
    }
  },
  {
    field: 'menu_type',
    label: t('system.menu.menuType'),
    component: 'RadioButton',
    value: 0,
    colProps: {
      span: 12
    },
    componentProps: {
      options: [
        {
          label: t('system.menu.directory'),
          value: 'M'
        },
        {
          label: t('common.menu'),
          value: 'C'
        },
        {
          label: t('system.menu.button'),
          value: 'F'
        },
        {
          label: t('system.menu.tab'),
          value: 'T'
        }
      ],
      disabled: true
    }
  },
  {
    field: 'icon',
    label: t('router.icon'),
    colProps: {
      span: 12
    },
    formItemProps: {
      slots: {
        default: (formData: any) =>
          h(IconPicker, {
            modelValue: formData.icon,
            'onUpdate:modelValue': (val: string) => {
              formData.icon = val
            }
          })
      }
    }
  },
  {
    field: 'path',
    label: t('system.menu.router'),
    colProps: colsSpan,
    component: 'Input',
    formItemProps: {
      rules: [required()]
    }
  },
  {
    field: 'component',
    label: t('system.menu.componentPath'),
    component: 'Input',
    colProps: colsSpan,
    value: '#',
    componentProps: {
      disabled: true,
      placeholder: t('system.menu.directoryRemark'),
      on: {
        change: (val: string) => {
          cacheComponent.value = val
        }
      }
    }
  },
  {
    field: 'api_id',
    label: t('system.menu.api'),
    component: 'Select',
    colProps: colsSpan,
    componentProps: {
      multiple: true,
      collapseTags: true
    },
    formItemProps: {
      slots: {
        default: (formdata) => {
          return h(SelectHeader, {
            options: apiList,
            maxCollapseTags: 2,
            modelValue: formdata.api_id,
            onSelectValues: async (value: any) => {
              setValues({ api_id: value })
            }
          })
        }
      }
    }
  },
  {
    field: 'visible',
    label: t('system.menu.visible'),
    component: 'Select',
    colProps: {
      span: 8
    },
    componentProps: {
      options: [
        { label: t('system.menu.show'), value: 1 },
        { label: t('system.menu.hide'), value: 2 },
        { label: t('system.menu.supervisor'), value: 3 }
      ],
      clearable: false
    },
    formItemProps: {
      rules: [required()]
    }
  },
  {
    field: 'status',
    label: t('userDemo.status'),
    component: 'Switch',
    colProps: {
      span: 8
    },
    componentProps: {
      activeValue: 1,
      inactiveValue: 2
    }
  },
  {
    field: 'is_cache',
    label: t('system.menu.isCache'),
    component: 'Switch',
    colProps: {
      span: 8
    },
    componentProps: {
      activeValue: 1,
      inactiveValue: 2
    }
  }
])

const rules = reactive({
  component: [required()],
  path: [required()],
  'meta.title': [required()]
})

const { formRegister, formMethods } = useForm()
const { setValues, getFormData, getElFormExpose, setSchema } = formMethods

const cacheComponent = ref('')
async function fetchMenuDetail(menuId: any) {
  const res = await fetchMenuDetailApi({ id: menuId })
  if (res.code === 0 && res.data) {
    const currentRow = res.data
    cacheComponent.value = currentRow.type === 1 ? currentRow.component : ''
    await setSchema([
      {
        field: 'component',
        path: 'remove',
        value: currentRow.menu_type === 'T' || currentRow.menu_type === 'F'
      },
      {
        field: 'component',
        path: 'componentProps.disabled',
        value: currentRow.menu_type !== 'C'
      },
      {
        field: 'path',
        path: 'remove',
        value: currentRow.menu_type === 'T' || currentRow.menu_type === 'F'
      }
    ])
    await setValues(currentRow)
  }
}

export async function menuWrite(callback: () => void, row?: any, menuData?: any) {
  apiList.value = []
  menuList.value = []
  DialogApi({
    title: t(row ? 'exampleDemo.edit' : 'exampleDemo.add'),
    fullscreen: true,
    showConfirmButton: true,
    body: () => h(Form, { onRegister: formRegister, schema: schema, rules: rules }),
    confirmAction: async (done: () => void) => {
      const elForm = await getElFormExpose()
      const valid = await elForm?.validate().catch((err) => {
        console.log(err)
      })
      if (valid) {
        const formData = await getFormData()
        const res: IResponse<any> = await menuModifyApi(formData)
        if (res?.code == 0) {
          ElMessage.success(t('common.operaSuccess'))
          if (row.menu_id == formData.menu_id) {
            Object.assign(row, formData)
          } else {
            if (row?.prent_id == formData?.prent_id) {
              row.children.map((item: any) => {
                if (item.menu_id == formData.menu_id) Object.assign(row, formData)
              })
            }
          }
          console.log(res, row)
          // callback()
          done()
        }
      }
    }
  })
  if (row) {
    await getApiList()
    menuList.value = menuData
    menuList.value.push({ menu_id: 0, title: t('system.menu.topMenu') })
    await setSchema([
      {
        field: 'parent_id',
        path: 'componentProps.data',
        value: menuData
      }
    ])
    await fetchMenuDetail(row.menu_id)
  }
}
