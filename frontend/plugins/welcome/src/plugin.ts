import { createPlugin } from '@backstage/core';
import CreateDeposit from './components/RecordDeposit';
import ShowDeposit from './components/DepositTable';
import SignIn from './components/SignIn';
import CreateRoomdetail from './components/RoomDetails';
import WelcomePage from './components/WelcomePage';
import CreateRepairinvoice from './components/RecordRepairinvoice';
import RepairinvoiceTablet from './components/RepairinvoiceTable';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
	  router.registerRoute('/DormEmployee', WelcomePage);
    router.registerRoute('/DepositTable', ShowDeposit);
    router.registerRoute('/RecordDeposit', CreateDeposit);
    router.registerRoute('/RoomDetails', CreateRoomdetail);
    router.registerRoute('/RecordRepairinvoice', CreateRepairinvoice);
    router.registerRoute('/RepairinvoiceTable', RepairinvoiceTablet );

  },
});
