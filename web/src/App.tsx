import { registerIcons, Theme, ThemeProvider } from '@fluentui/react';
import { AscendingIcon, Blocked2Icon, CancelIcon, CheckMarkIcon, ChevronDownIcon, ChevronDownSmallIcon, ChevronLeftIcon, ChevronRightIcon, ChevronUpSmallIcon, CircleAdditionIcon, CircleAdditionSolidIcon, CircleRingIcon, ClearIcon, ClipboardListIcon, CodeIcon, ColorSolidIcon, CompletedIcon, CopyIcon, DatabaseIcon, DataManagementSettingsIcon, DeleteIcon, DescendingIcon, DoubleChevronDownIcon, DoubleChevronUpIcon, EditIcon, EmbedIcon, Emoji2Icon, ErrorBadgeIcon, FilterIcon, FilterSolidIcon, FunnelChartIcon, GroupListIcon, HideIcon, HomeIcon, InfoIcon, InstallationIcon, LocaleLanguageIcon, LocationIcon, MapLayersIcon, MoreIcon, MoreVerticalIcon, PermissionsIcon, PlugConnectedIcon, PlugDisconnectedIcon, ProductListIcon, PublishContentIcon, QueryListIcon, RedEyeIcon, RefreshIcon, RemoveFromShoppingListIcon, RevToggleKeyIcon, SadIcon, SaveIcon, SearchDataIcon, SearchIcon, ServerProcessesIcon, SettingsIcon, SkypeCircleCheckIcon, StatusCircleCheckmarkIcon, StatusCircleErrorXIcon, StatusErrorFullIcon, SyncIcon, UploadIcon, ViewIcon } from '@fluentui/react-icons';
import React, { useEffect, useState } from 'react';
import { I18nextProvider } from 'react-i18next';
import './App.css';
import { ErrorMessageBar } from './components/common/ErrorMessageBar';
import { Loading } from './components/common/Loading';
import { MainTab } from './components/MainTab';
import i18n from './i18n';
import { Config, getConfig, setConfig } from './services/config.service';
import { themes } from './theme';

// icons
registerIcons({
  icons: {
    circleAdditionSolid: <CircleAdditionSolidIcon />,
    circleAddition: <CircleAdditionIcon />,
    publishcontent: <PublishContentIcon />,
    installation: <InstallationIcon />,
    chevrondown: <ChevronDownIcon />,
    chevronright: <ChevronRightIcon />,
    mapLayers: <MapLayersIcon />,
    database: <DatabaseIcon />,
    permissions: <PermissionsIcon />,
    moreVertical: <MoreVerticalIcon />,
    clipboardList: <ClipboardListIcon />,
    embed: <EmbedIcon />,
    refresh: <RefreshIcon />,
    copy: <CopyIcon />,
    edit: <EditIcon />,
    plugDisconnected: <PlugDisconnectedIcon />,
    plugConnected: <PlugConnectedIcon />,
    filter: <FilterIcon />,
    filterSolid: <FilterSolidIcon />,
    funnelChart: <FunnelChartIcon />,
    more: <MoreIcon />,
    statusErrorFull: <StatusErrorFullIcon />,
    statusCircleErrorX: <StatusCircleErrorXIcon />,
    cancel: <CancelIcon />,
    chevrondownsmall: <ChevronDownSmallIcon />,
    chevronupsmall: <ChevronUpSmallIcon />,
    chevronleft: <ChevronLeftIcon />,
    redeye: <RedEyeIcon />,
    hide: <HideIcon />,
    doublechevrondown: <DoubleChevronDownIcon />,
    doublechevronup: <DoubleChevronUpIcon />,
    blocked2: <Blocked2Icon />,
    errorbadge: <ErrorBadgeIcon />,
    clear: <ClearIcon />,
    completed: <CompletedIcon />,
    delete: <DeleteIcon />,
    removeFromShoppingList: <RemoveFromShoppingListIcon />,
    skypeCircleCheck: <SkypeCircleCheckIcon />,
    save: <SaveIcon />,
    revToggleKey: <RevToggleKeyIcon />,
    statuscirclecheckmark: <StatusCircleCheckmarkIcon />,
    circlering: <CircleRingIcon />,
    checkmark: <CheckMarkIcon />,
    grouplist: <GroupListIcon />,
    querylist: <QueryListIcon />,
    ascending: <AscendingIcon />,
    descending: <DescendingIcon />,
    view: <ViewIcon />,
    settings: <SettingsIcon />,
    colorsolid: <ColorSolidIcon />,
    localeLanguage: <LocaleLanguageIcon />,
    home: <HomeIcon />,
    search: <SearchIcon />,
    searchdata: <SearchDataIcon />,
    syncicon: <SyncIcon />,
    info: <InfoIcon />,
    ServerProcesses: <ServerProcessesIcon />,
    DataManagementSettings: <DataManagementSettingsIcon />,
    Code: <CodeIcon />,
    ProductList: <ProductListIcon />,
    Upload: <UploadIcon />,
    Sad: <SadIcon />,
    Emoji2: <Emoji2Icon />,
    Location: <LocationIcon />
  }
})


function App() {
  const [_config, _setConfig] = useState<Config>({ theme: 'light', language: 'en_US', port: 8080 }),
    [theme, setTheme] = useState<Theme>(themes.light),
    [error, setError] = useState<Error>(),
    [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    setLoading(true);
    getConfig().then((ret) => {
      if (!ret) return;
      _setConfig({ ...ret });
      setTheme(themes[ret.theme]);
    })
      .catch(setError)
      .finally(() => {
        setLoading(false);
      })
  }, []);

  useEffect(() => {
    _config.language && i18n.changeLanguage(_config.language);
  }, [_config.language])

  const handleChangeTheme = (theme: string) => {
    changeConfig({ ..._config, theme });
    setTheme(themes[theme]);
  }

  const handleChnageLanguage = (language: string) => {
    changeConfig({ ..._config, language });
  }

  const handleChangePort = (port: number) => {
    changeConfig({ ..._config, port })
  }

  const changeConfig = (conf: Config) => {
    setLoading(true);
    setConfig(conf)
      .then(ret => {
        _setConfig(ret)
      })
      .catch(setError)
      .finally(() => {
        setLoading(false);
      })
  }

  return (<>
    <ThemeProvider theme={theme} style={{ height: '100%' }}>
      <I18nextProvider i18n={i18n}>
        {/* loading */}
        <Loading loading={loading} />
        {/* error */}
        <ErrorMessageBar error={error}></ErrorMessageBar>
        {/* mainTab */}
        <MainTab config={_config}
          onChangeLanguage={handleChnageLanguage}
          onChangeTheme={handleChangeTheme}
          onChangePort={handleChangePort}
        ></MainTab>
      </I18nextProvider>
    </ThemeProvider>
  </>);
}

export default App;
