export const useMenuWrite = () => {
  const getFormSchema = (menuOptions = []) => {
    return [
      {
        field: 'parentId', label: '上级菜单', component: 'treeSelect',
        options: menuOptions,
        treeProps: { label: 'name', value: 'id', children: 'children' },
        checkStrictly: true,
      },
      {
        field: 'type', label: '菜单类型', component: 'radio',
        options: [
          { value: 1, label: '目录' },
          { value: 2, label: '菜单' },
          { value: 3, label: '按钮' },
        ],
      },
      { field: 'name', label: '菜单名称', component: 'input', rules: [{ required: true, message: '请输入菜单名称' }] },
      { field: 'icon', label: '图标', component: 'input', placeholder: '例: User, Setting' },
      { field: 'path', label: '路由路径', component: 'input', vIf: 'type === 2' },
      { field: 'component', label: '组件路径', component: 'input', vIf: 'type === 2' },
      { field: 'permission', label: '权限标识', component: 'input' },
      { field: 'sort', label: '排序', component: 'number' },
      { field: 'status', label: '状态', component: 'switch' },
    ]
  }

  return { getFormSchema }
}
