jQuery(($) => {

    $(document).ready(() => {
        autosize($('textarea'));
        scrollUp();
        popupInit();
        sidebarInit();
        formValidation();
        geoMapInit();
        selectInit();
        inputDateInit();
        dropdownInit();
        ISTabs();
        ISMenu();
        copyField();
        commentsInit();
    });

    const scrollUp = () => {
        const scrollUpBtn = $('.scroll-up');
        let currentScroll;

        $(window).on('scroll resize', () => {
            currentScroll = $(window).scrollTop();

            if (currentScroll > 200) {
                scrollUpBtn.show();
            } else {
                scrollUpBtn.hide();
            }
        });

        scrollUpBtn.on('click', (e) => {
            e.preventDefault();
            $('html, body').animate({
                scrollTop: 0
            }, 500);
        });
    };

    const popupInit = () => {
        const popup = $('.popup');
        const popupWindow = $('.popup__block');
        const popupCloseBtn = $('.popup__close');
        const popupLink = $('.popupOpen');

        $(document).on('mouseup', (e) => {
            if (!popupWindow.is(e.target) && popupWindow.has(e.target).length === 0) {
                popupClose();
            }
        });

        popupLink.on('click', function (e) {
            let id = $(this).attr('href');
            console.log(this);

            if ($(id).length === 0) return;

            e.preventDefault();
            popupOpen(id);
        });

        popupCloseBtn.on('click', (e) => {
           e.preventDefault();
           popupClose();
        });

        function popupOpen(el) {
            $(el).fadeIn();
            $('body').addClass('no-scroll');
        }

        function popupClose() {
            $('body').removeClass('no-scroll');
            popup.hide();
        }
    };

    const sidebarInit = () => {
        const sidebar = $('.sidebar');
        const sidebarToggleBtn = $('.sidebar__toggle');

        sidebarToggleBtn.on('click', function (e) {
            e.preventDefault();
            let title = $(this).attr('title');
            let titleAlt = $(this).attr('data-title-alt');

            $(this).attr('title', titleAlt);
            $(this).attr('data-title-alt', title);

            sidebar.toggleClass('sidebar_opened');
        });
    }

    const formValidation = () => {
        $('.form').each(function () {
            $(this).validate();
        })
    };

    const geoMapInit = () => {
        const map = $('.geo__map svg');
        const region = $('.geo__region');
        const mapInfo = $('.geo__map-info');

        region.on('click', function () {
            let regionNumber = $(this).data('region');

            if ($(this).hasClass('geo__region_active')) return;

            hoverRegion(regionNumber);
        });

        map.find('path').on('mouseenter', function (e) {
            let regionNumber = $(e.target).data('region');
            if (!regionNumber) return;
            hoverRegion(regionNumber);
        });

        map.find('path').on('click', function () {
            let href = $(this).data('href');
            if (!href) return;
            window.location.href = href;
        });

        function hoverRegion(regionNumber) {
            let el = $('.geo__region[data-region="' + regionNumber + '"]').eq(0);
            let name = el.find('.geo__region-name').text();
            let num = +el.find('.geo__region-count').text();

            region.removeClass('geo__region_active');
            el.addClass('geo__region_active');

            map.find('path').removeClass('active');
            map.find('path[data-region="' + regionNumber + '"]').addClass('active').appendTo(map.find('path').parent());

            mapInfo.hide();
            mapInfo.find('.geo__map-num').text(num);
            if (num < 1) {
                mapInfo.find('.geo__map-num').addClass('geo__map-num_empty');
            } else {
                mapInfo.find('.geo__map-num').removeClass('geo__map-num_empty');
            }
            mapInfo.find('.geo__map-name span').text(name);
            mapInfo.attr('data-region', regionNumber).fadeIn();
        }
    };

    const selectInit = () => {
        $('select').on('change', function () {
           if ($(this).val() === '') {
               $(this).removeClass('filled');
           } else {
               $(this).addClass('filled');
           }
        });
    };

    const inputDateInit = () => {
        const settings = {
            showOtherMonths: true,
            selectOtherMonths: true,
            dateFormat: 'dd.mm.yy',
            prevText: '',
            nextText: ''
        };

        $('input[type="date"]').each(function () {
            $(this).attr('type', 'text').datepicker(settings);
        });

        const exportFrom = $('#export-from')
            .datepicker(settings)
            .on('change', function () {
                console.log(getDate(this))
                exportTo.datepicker('option', 'minDate', getDate(this));
            });
        const exportTo = $('#export-to')
            .datepicker(settings)
            .on('change', function () {
                exportFrom.datepicker('option', 'maxDate', getDate(this));
            });

        function getDate(element) {
            let date;
            try {
                date = $.datepicker.parseDate( settings.dateFormat, element.value );
            } catch(error) {
                date = null;
            }
            return date;
        }
    };

    const dropdownInit = () => {
        $('.dropdown .dropdown__title').on('click', function () {
            $(this).parent().toggleClass('dropdown_active').find('.dropdown__content').slideToggle();
        });
    };

    const ISTabs = () => {
        const tabs = $('.is-tabs');
        const tab = tabs.find('.tabs__link');
        const tabContent = $('.is-tab-content');
        const tabActiveClass = 'tabs__link_active';

        tab.on('click', function (e) {
            e.preventDefault();
            if ($(this).hasClass(tabActiveClass)) return;

            tab.removeClass(tabActiveClass);
            $(this).addClass(tabActiveClass);

            let id = $(this).data('id');

            tabContent.hide()
            $('.is-tab-content[data-id="' + id + '"]').fadeIn();
        });
    };

    const ISMenu = () => {
        const menu = $('.is-menu');
        const menuItem = menu.find('.is-menu__item');
        const submenuItem = menu.find('.is-menu__item-submenu li');
        const menuContent = $('.is-menu-content');
        const menuToggleButton = $('.is-menu__toggle');
        const menuBackButton = $('.is-menu__back');
        const menuItemActiveClass = 'is-menu__item_active';

        menuToggleButton.on('click', function (e) {
            e.preventDefault();
            $(this).parents('.is-menu').toggleClass('is-menu_closed');
        });

        menuBackButton.on('click', () => {
            $('.is-menu_another').hide();
            $('.is-menu_main').fadeIn();
        });

        menuItem.on('click', function (e) {
            e.preventDefault();
            if ($(this).hasClass('is-menu__item_has-submenu')) return;

            let id = $(this).data('id');

            if ($(this).hasClass('is-menu__item_another-menu')) {
                menu.hide();
                $('.is-menu[data-id="' + id + '"]').fadeIn();
                return;
            }

            addActiveClass($(this), menuItemActiveClass);
            showContent(id);
        });

        submenuItem.on('click', function (e) {
            e.preventDefault();

            let id = $(this).data('id');

            addActiveClass($(this), 'active');
            showContent(id);
        });

        function addActiveClass(el, activeClass) {
            menuItem.removeClass(menuItemActiveClass)
            submenuItem.removeClass('active');
            el.addClass(activeClass);
        }

        function showContent(id) {
            menuContent.hide();
            $('.is-menu-content[data-id="' + id + '"]').fadeIn();
        }
    };

    const copyField = () => {
        $('.copy-field__btn').on('click', function (e) {
            e.preventDefault();
            let parent = $(this).parents('.copy-field');
            let content = parent.find('.copy-field__content');

            parent.append('<textarea id="tempTextarea">' + content.text().trim() + '</textarea>');
            $('#tempTextarea').select();

            document.execCommand('copy');

            $('#tempTextarea').remove();

            parent.addClass('copy-field_copied');

            setTimeout(() => {
                parent.removeClass('copy-field_copied');
            }, 2000);
        });
    }

    const commentsInit = () => {
        const commentsEditForm = $('.comments__edit-form');
        const commentsDeleteForm = $('.comments__delete-form');
        const textarea = commentsEditForm.find('textarea');

        autosize.update(textarea);

        $(document).on('click', '.comments-button', () => {
            $('.comments').addClass('comments_active');
        });

        $(document).on('click', '.comments__close', () => {
            $('.comments').removeClass('comments_active');
        });

        $(document).on('click', '.comment__edit', function () {
            let comment = $(this).parents('.comment');
            let commentId = comment.attr('id');
            let commentContent = comment.find('.comment__content').text().trim();

            if (!commentId) return;

            commentsEditForm.find('input[name="comment-id"]').val(commentId);
            commentsEditForm.find('textarea').val(commentContent);
            setTimeout(() => {
                autosize.update(textarea);
            }, 1);
            commentsEditForm.hide().appendTo(comment).fadeIn();
        });

        $(document).on('click', '.comment__delete', function () {
            let comment = $(this).parents('.comment');
            let commentId = comment.attr('id');

            if (!commentId) return;

            commentsDeleteForm.find('.comments__delete-form-submit').attr('data-comment_id', commentId);
            commentsDeleteForm.hide().appendTo(comment).fadeIn();
        });

        $(document).on('click', '.comments__close-icon, .comments__delete-form-close', function (e) {
            e.preventDefault();

            if ($(this).hasClass('comments__edit-form-close')) {
                commentsEditForm.fadeOut(() => {
                    commentsEditForm.appendTo('.comments');
                });
            } else if ($(this).hasClass('comments__delete-form-close')) {
                commentsDeleteForm.fadeOut(() => {
                    commentsDeleteForm.appendTo('.comments');
                });
            }
        });

        $(document).on('submit', '.comments__edit-form', (e) => {
            e.preventDefault();
            let commentId = commentsEditForm.find('input[name="comment-id"]').val(),
                commentContent = commentsEditForm.find('textarea[name="comment"]').val(),
                comment = $('#' + commentId),
                data = {
                    action: 'comment_edit',
                    commentId,
                    commentContent
                };

            /* AJAX Handler */

            // on success
            comment.hide().find('.comment__content').text(commentContent);
            commentsEditForm.hide().appendTo('.comments');
            comment.fadeIn();
        });

        $(document).on('click', '.comments__delete-form-submit', function (e) {
            e.preventDefault();

            let commentId = $(this).attr('data-comment_id'),
                comment = $('#' + commentId),
                data = {
                    action: 'comment_delete',
                    commentId
                };

            /* AJAX Handler */

            // on success
            commentsEditForm.hide().appendTo('.comments');
            commentsDeleteForm.hide().appendTo('.comments');
            comment.remove();
        });
    };

});